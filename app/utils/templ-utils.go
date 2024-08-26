package utils

import "strings"

func CreateClassName(classNames ...string) string {
	return strings.Join(classNames, " ")
}

func ReplaceColorPlaceholder(class, color string) string {
	return strings.ReplaceAll(class, "{color}", color)
}

var VariantClassNames = map[string]map[string]string{
	"color": {
		"black":		"bg-black hover:bg-black text-white",
		"white":		"bg-white hover:bg-white text-black",
		"slate":		"bg-slate-600 hover:bg-slate-500 text-white",
		"slate-dark":		"bg-slate-900 hover:bg-slate-800 text-white",
		"gray":		"bg-gray-600 hover:bg-gray-500 text-white",
		"zinc":		"bg-zinc-600 hover:bg-zinc-500 text-white",
		"neutral":		"bg-neutral-600 hover:bg-neutral-500 text-white",
		"stone":		"bg-stone-600 hover:bg-stone-500 text-white",
		"stone-dark":		"bg-stone-900 hover:bg-stone-800 text-white",
		"red":		"bg-red-600 hover:bg-red-500 text-white",
		"orange":		"bg-orange-600 hover:bg-orange-500 text-white",
		"amber":		"bg-amber-600 hover:bg-amber-500 text-white",
		"yellow":		"bg-yellow-600 hover:bg-yellow-500 text-white",
		"lime":		"bg-lime-600 hover:bg-lime-500 text-white",
		"green":		"bg-green-600 hover:bg-green-500 text-white",
		"emerald":		"bg-emerald-600 hover:bg-emerald-500 text-white",
		"teal":		"bg-teal-600 hover:bg-teal-500 text-white",
		"cyan":		"bg-cyan-600 hover:bg-cyan-500 text-white",
		"sky":		"bg-sky-600 hover:bg-sky-500 text-white",
		"blue":		"bg-blue-600 hover:bg-blue-500 text-white",
		"indigo":		"bg-indigo-600 hover:bg-indigo-500 text-white",
		"violet":		"bg-violet-600 hover:bg-violet-500 text-white",
		"purple":		"bg-purple-600 hover:bg-purple-500 text-white",
		"fuchsia":		"bg-fuchsia-600 hover:bg-fuchsia-500 text-white",
		"pink":		"bg-pink-600 hover:bg-pink-500 text-white",
		"rose":		"bg-rose-600 hover:bg-rose-500 text-white",
	},
	"outline": {
		"black":   "bg-transparent border border-black text-black hover:bg-black hover:text-white",
		"white":   "bg-transparent border border-white text-white hover:bg-white hover:text-black",
		"slate":   "bg-transparent border border-slate-600 text-slate-600 hover:bg-slate-600 hover:text-white",
		"slate-dark":   "bg-transparent border border-slate-900 text-slate-900 hover:bg-slate-900 hover:text-white",
		"gray":    "bg-transparent border border-gray-600 text-gray-600 hover:bg-gray-600 hover:text-white",
		"zinc":    "bg-transparent border border-zinc-600 text-zinc-600 hover:bg-zinc-600 hover:text-white",
		"neutral": "bg-transparent border border-neutral-600 text-neutral-600 hover:bg-neutral-600 hover:text-white",
		"stone":   "bg-transparent border border-stone-600 text-stone-600 hover:bg-stone-600 hover:text-white",
		"stone-dark":   "bg-transparent border border-stone-900 text-stone-900 hover:bg-stone-900 hover:text-white",
		"red":     "bg-transparent border border-red-600 text-red-600 hover:bg-red-600 hover:text-white",
		"orange":  "bg-transparent border border-orange-600 text-orange-600 hover:bg-orange-600 hover:text-white",
		"amber":   "bg-transparent border border-amber-600 text-amber-600 hover:bg-amber-600 hover:text-white",
		"yellow":  "bg-transparent border border-yellow-600 text-yellow-600 hover:bg-yellow-50600 hover:text-white",
		"lime":    "bg-transparent border border-lime-600 text-lime-600 hover:bg-lime-600 hover:text-white",
		"green":   "bg-transparent border border-green-600 text-green-600 hover:bg-green-600 hover:text-white",
		"emerald": "bg-transparent border border-emerald-600 text-emerald-600 hover:bg-emerald-600 hover:text-white",
		"teal":    "bg-transparent border border-teal-600 text-teal-600 hover:bg-teal-600 hover:text-white",
		"cyan":    "bg-transparent border border-cyan-600 text-cyan-600 hover:bg-cyan-50600 hover:text-white",
		"sky":     "bg-transparent border border-sky-600 text-sky-600 hover:bg-sky-600 hover:text-white",
		"blue":    "bg-transparent border border-blue-600 text-blue-600 hover:bg-blue-600 hover:text-white",
		"indigo":  "bg-transparent border border-indigo-600 text-indigo-600 hover:bg-indigo-600 hover:text-white",
		"violet":  "bg-transparent border border-violet-600 text-violet-600 hover:bg-violet-600 hover:text-white",
		"purple":  "bg-transparent border border-purple-600 text-purple-600 hover:bg-purple-600 hover:text-white",
		"fuchsia": "bg-transparent border border-fuchsia-600 text-fuchsia-600 hover:bg-fuchsia-600 hover:text-white",
		"pink":    "bg-transparent border border-pink-600 text-pink-600 hover:bg-pink-600 hover:text-white",
		"rose":    "bg-transparent border border-rose-600 text-rose-600 hover:bg-rose-600 hover:text-white",
	},
	"link": {
		"black":   "bg-transparent hover:bg-transparent text-black hover:text-black underline",
		"white":   "bg-transparent hover:bg-transparent text-white hover:text-white underline",
		"slate":   "bg-transparent hover:bg-transparent text-slate-600 hover:text-slate-500 underline",
		"slate-dark":   "bg-transparent hover:bg-transparent text-slate-900 hover:text-slate-500 underline",
		"gray":    "bg-transparent hover:bg-transparent text-gray-600 hover:text-gray-500 underline",
		"zinc":    "bg-transparent hover:bg-transparent text-zinc-600 hover:text-zinc-500 underline",
		"neutral": "bg-transparent hover:bg-transparent text-neutral-600 hover:text-neutral-500 underline",
		"stone":   "bg-transparent hover:bg-transparent text-stone-600 hover:text-stone-500 underline",
		"red":     "bg-transparent hover:bg-transparent text-red-600 hover:text-red-500 underline",
		"orange":  "bg-transparent hover:bg-transparent text-orange-600 hover:text-orange-500 underline",
		"amber":   "bg-transparent hover:bg-transparent text-amber-600 hover:text-amber-500 underline",
		"yellow":  "bg-transparent hover:bg-transparent text-yellow-600 hover:text-yellow-500 underline",
		"lime":    "bg-transparent hover:bg-transparent text-lime-600 hover:text-lime-500 underline",
		"green":   "bg-transparent hover:bg-transparent text-green-600 hover:text-green-500 underline",
		"emerald": "bg-transparent hover:bg-transparent text-emerald-600 hover:text-emerald-500 underline",
		"teal":    "bg-transparent hover:bg-transparent text-teal-600 hover:text-teal-500 underline",
		"cyan":    "bg-transparent hover:bg-transparent text-cyan-600 hover:text-cyan-500 underline",
		"sky":     "bg-transparent hover:bg-transparent text-sky-600 hover:text-sky-500 underline",
		"blue":    "bg-transparent hover:bg-transparent text-blue-600 hover:text-blue-500 underline",
		"indigo":  "bg-transparent hover:bg-transparent text-indigo-600 hover:text-indigo-500 underline",
		"violet":  "bg-transparent hover:bg-transparent text-violet-600 hover:text-violet-500 underline",
		"purple":  "bg-transparent hover:bg-transparent text-purple-600 hover:text-purple-500 underline",
		"fuchsia": "bg-transparent hover:bg-transparent text-fuchsia-600 hover:text-fuchsia-500 underline",
		"pink":    "bg-transparent hover:bg-transparent text-pink-600 hover:text-pink-500 underline",
		"rose":    "bg-transparent hover:bg-transparent text-rose-600 hover:text-rose-500 underline",	
	},
	"unstyled": {
		"black":   "bg-transparent hover:bg-transparent",
		"white":   "bg-transparent hover:bg-transparent",
		"slate":   "bg-transparent hover:bg-transparent",
		"gray":    "bg-transparent hover:bg-transparent",
		"zinc":    "bg-transparent hover:bg-transparent",
		"neutral": "bg-transparent hover:bg-transparent",
		"stone":   "bg-transparent hover:bg-transparent",
		"red":     "bg-transparent hover:bg-transparent",
		"orange":  "bg-transparent hover:bg-transparent",
		"amber":   "bg-transparent hover:bg-transparent",
		"yellow":  "bg-transparent hover:bg-transparent",
		"lime":    "bg-transparent hover:bg-transparent",
		"green":   "bg-transparent hover:bg-transparent",
		"emerald": "bg-transparent hover:bg-transparent",
		"teal":    "bg-transparent hover:bg-transparent",
		"cyan":    "bg-transparent hover:bg-transparent",
		"sky":     "bg-transparent hover:bg-transparent",
		"blue":    "bg-transparent hover:bg-transparent",
		"indigo":  "bg-transparent hover:bg-transparent",
		"violet":  "bg-transparent hover:bg-transparent",
		"purple":  "bg-transparent hover:bg-transparent",
		"fuchsia": "bg-transparent hover:bg-transparent",
		"pink":    "bg-transparent hover:bg-transparent",
		"rose":    "bg-transparent hover:bg-transparent",	
	},
}

var CornerClassNames = map[string]string{
	"rounded": "rounded-md",
	"sharp":   "rounded",
	"soft":    "rounded-lg",
	"none":    "",
}

var SizeClassNames = map[string]string{
	"sm": "px-2 py-1 text-sm",
	"md": "px-3 py-2 text-sm",
	"lg": "px-4 py-2 text-base",
	"xl": "px-6 py-3 text-base",
	"2xl": "px-8 py-3 text-base",
	"3xl": "px-10 py-4 text-base",
	"4xl": "px-12 py-6 text-lg",
	"5xl": "px-16 py-8 text-lg",
	"6xl": "px-20 py-8 text-xl",
}

func GetVariantClassName(variant, color string) string {
	if variantMap, ok := VariantClassNames[variant]; ok {
		if className, ok := variantMap[color]; ok {
			return className
		}
	}
	return VariantClassNames["color"]["blue"] // default
}

func GetCornerClassName(corners string) string {
	if className, ok := CornerClassNames[corners]; ok {
		return className
	}
	return CornerClassNames["rounded"] // default
}

func GetSizeClassName(size string) string {
	if className, ok := SizeClassNames[size]; ok {
		return className
	}
	return SizeClassNames["md"] // default
}

func GetFocusClassName(color string) string {
	return "focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-" + color + "-600"
}