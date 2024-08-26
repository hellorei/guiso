package logger

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
)

var (
	once   sync.Once
	log    zerolog.Logger
	IsDev  bool
)

// Color codes
const (
	ColorReset  = "\033[0m"
	ColorPurple = "\033[38;2;203;166;247m"
	ColorRosewater = "\033[38;2;245;224;220m"
	ColorFlamingo = "\033[38;2;242;205;205m"
	ColorPink = "\033[38;2;245;194;231m"
	ColorMauve = "\033[38;2;203;166;247m"
	ColorRed = "\033[38;2;243;139;168m"
	ColorMaroon = "\033[38;2;235;160;172m"
	ColorPeach = "\033[38;2;250;179;135m"
	ColorYellow = "\033[38;2;249;226;175m"
	ColorGreen = "\033[38;2;166;227;161m"
	ColorTeal = "\033[38;2;148;226;213m"
	ColorSky = "\033[38;2;137;220;235m"
	ColorSapphire = "\033[38;2;116;199;236m"
	ColorBlue = "\033[38;2;137;180;250m"
	ColorLavender = "\033[38;2;180;190;254m"
	ColorLavender_light = "\033[38;2;180;190;254m"
	ColorLavender_dark = "\033[38;2;127;132;156m"
	ColorText = "\033[38;2;205;214;244m"
	ColorSubtext1 = "\033[38;2;186;194;222m"
	ColorSubtext0 = "\033[38;2;166;173;200m"
	ColorOverlay2 = "\033[38;2;147;153;178m"
	ColorOverlay1 = "\033[38;2;127;132;156m"
	ColorOverlay0 = "\033[38;2;108;112;134m"
	ColorSurface2 = "\033[38;2;88;91;112m"
	ColorSurface1 = "\033[38;2;69;71;90m"
	ColorSurface0 = "\033[38;2;69;71;90m"
	ColorBase = "\033[38;2;30;30;46m"
	ColorMantle = "\033[38;2;24;24;37m"
	ColorCrust = "\033[38;2;17;17;27m"
)

func init() {
	IsDev = os.Getenv("ENV") != "production"
	initLogger()
}

func initLogger() {
	once.Do(func() {
		var output io.Writer
		if IsDev && isatty.IsTerminal(os.Stdout.Fd()) {
			output = zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: "[15:04:05.000]",
				NoColor:    false,
				FormatLevel: func(i interface{}) string {
					var l string
					if ll, ok := i.(string); ok {
							switch ll {
							case "debug":
									l = Colorize(" DEBUG", ColorGreen)
							case "info":
									l = Colorize(" INFO ", ColorOverlay2)
							case "warn":
									l = Colorize(" WARN ", ColorYellow)
							case "error":
									l = Colorize(" ERROR", ColorRed)
							case "fatal":
									l = Colorize(" FATAL", ColorRed)
							case "panic":
									l = Colorize(" PANIC", ColorPurple)
							default:
									l = Colorize(ll, ColorGreen)
							}
					}
					return fmt.Sprintf("%-6s", l)
				},
				FormatMessage: func(i interface{}) string {
					return Colorize(fmt.Sprintf("%s", i), ColorSurface2)
				},
				FormatFieldName: func(i interface{}) string {
					return Colorize(fmt.Sprintf("%s=", i), ColorYellow)
				},
				FormatFieldValue: func(i interface{}) string {
					return Colorize(fmt.Sprintf("%s", i), ColorGreen)
				},
			}
		} else {
			output = os.Stdout
		}

		logLevel := zerolog.InfoLevel
		if IsDev {
			logLevel = zerolog.DebugLevel
		}

		zerolog.SetGlobalLevel(logLevel)
		log = zerolog.New(output).With().Timestamp().Logger()
	})
}

// GetLogger returns the configured logger
func GetLogger() zerolog.Logger {
	return log
}

// Colorize returns a colorized string
func Colorize(s, color string) string {
	return color + s + ColorReset
}

// Info logs an info message
func Info(msg string) {
	log.Info().Msg(msg)
}

// Debug logs a debug message
func Debug(msg string) {
	log.Debug().Msg(msg)
}

// Error logs an error message
func Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}

// Warn logs a warning message
func Warn(msg string) {
	log.Warn().Msg(msg)
}