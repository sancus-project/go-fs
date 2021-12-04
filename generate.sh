#!/bin/sh

set -eu
F="${0%.sh}_sh.go"
trap "rm -f '$F~'" EXIT
exec > "$F~"

cat <<EOT
package ${GOPACKAGE:-undefined}

//go:generate $0${1:+ $*}

// Code generated by $0. DO NOT EDIT.
EOT

generate_alias() {
	local t="$1" pkg="$2" n=
	shift 2

	if [ $# -gt 1 ]; then
		cat <<EOT

${t#*:} (
EOT
		for n; do
			cat <<EOT
	// $n is an alias of the $pkg.$n ${t#*:},
	// see https://pkg.go.dev/$pkg#$n for details.
	$n = ${pkg##*/}.$n
EOT
		done
		cat <<EOT
)
EOT
	elif [ $# -eq 1 ]; then
		n="$1"
		cat <<EOT

// $n is an alias of the $pkg.$n ${t#*:},
// see https://pkg.go.dev/$pkg#$n for details.
${t%:*} $n = ${pkg##*/}.$n
EOT
	fi
}

generate_const() {
	generate_alias const:constant "$@"
}

generate_types() {
	generate_alias type "$@"
}

generate_wrapped_alias() {
	local t="$1" pkg="$2" f="$3" n=
	shift 3

	if [ $# -gt 1 ]; then
		cat <<EOT

${t%:*} (
EOT
		for n; do
			cat <<EOT
	// $n is a wrapped alias of the $pkg.$n ${t#*:},
	// see https://pkg.go.dev/$pkg#$n for details.
	$n = $f(${pkg##*/}.$n)
EOT
		done
		cat <<EOT
)
EOT
	elif [ $# -eq 1 ]; then
		n="$1"
		cat <<EOT

// $n is a wrapped alias of the $pkg.$n ${t#*:},
// see https://pkg.go.dev/$pkg#$n for details.
${t%:*} $n = $f(${pkg##*/}.$n)
EOT
	fi
}

generate_wrapped_const() {
	generate_wrapped_alias "const:constant" "$@"
}

generate__proxy() {
	local pkg="$1" fn="$2"
	local vars="$3" types="$4" rets=
	local args= params= v=
	shift 4

	# prepare list of arguments
	for v in $vars; do
		args="${args:+$args, }$v @@T@@"
		params="${params:+$params, }$v"
	done

	# replace types one at the time
	for v in $types; do
		args="$(echo "$args" | sed -e "s/@@T@@/$v/")"
	done

	if [ $# -gt 1 ]; then
		rets=$(echo "$*" | sed -e "s/ /, /g")
		rets="($rets)"
	else
		rets="${1:-}"
	fi
	cat <<EOT

// $fn is a proxy function to $pkg.$fn(),
// see https://pkg.go.dev/$pkg#$fn for details.
func $fn($args)${rets:+ $rets} {
	return ${pkg##*/}.$fn($params)
}
EOT
}

generate_proxy() {
	local pkg="$1" fn="$2"
	local vars="$(echo "$3" | tr ',' ' ')"
	local types="$(echo "$4" | tr ',' ' ')"
	local rets="$(echo "${5:-}" | tr ',' ' ')"

	generate__proxy "$pkg" "$fn" "$vars" "$types" $rets
}

generate_proxies() {
	local pkg="$1" x=
	local fn= vars= types= rets=
	shift
	for x; do
		fn="$(echo "$x" | cut -d: -f1)"
		vars="$(echo "$x" | cut -d: -f2)"
		types="$(echo "$x" | cut -d: -f3)"
		rets="$(echo "$x" | cut -d: -f4)"

		generate_proxy "$pkg" "$fn" "$vars" "$types" "$rets"
	done
}

generate_done() {
	# reformat
	if ! gofmt -w -l -s "$F~"; then
		diff -u "$F" "$F~" >&2
		exit 1
	fi

	# diff
	if ! diff -u "$F" "$F~" >&2; then
		# replace if needed
		mv "$F~" "$F"
	else
		rm -f "$F~"
	fi
}
