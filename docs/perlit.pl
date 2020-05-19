#!/usr/bin/perl -p -i

BEGIN {
$p='/*
Intent:

@TODO - This has yet to be written.

Description:

@TODO - This has yet to be written.
	@TODO - This has yet to be written.
	@TODO - This has yet to be written.
*/
';

$f='/*
Intent:

@TODO - This has yet to be written.

Arguments:

%%ARGS%%

Helper examples:

- Default call
	`{{ %%FUNC%% %%FUNCARGS%% }}`
*/
';
}

if (/^package/) {
	s/^(package)/$p$1/;

} elsif (/^func (\w+)\((.*?)\)\s+(.*?)({| )/) {
	$fn = $1;
	$fa = $2;
	$fr = $3;
	$fp = $f;
	$fat = "";
	$fant = "";
	foreach $z (split(",", $fa)) {
		@zz = split(" ", $z);
		$fat .= sprintf("- \x60%s\x60 - Type: \x60%s\x60 - @TODO\n", $zz[0], $zz[1]);
		$fant .= sprintf(" <%s>", $zz[0]);
	}
	$fp =~ s/%%ARGS%%/$fat/;
	$fp =~ s/%%FUNCARGS%%/$fant/;
	if ($fr =~ /State/) {
		$fr = 'state';
	} else {
		$fr =~ s/\*//;
		$fr =~ s/^Helper//;
	}

	if ($fn =~ /^Helper/) {
		$fp =~ s/%%FUNC%%/\$$fr := $fn/g;
	} else {
		$fp =~ s/%%FUNC%%/\$state := $fn/g;
	}

	s/^(func)/$fp$1/;

#} elsif (/^func \(.*?\) (\w+)\(.*\*(.*?)({| )/) {
} elsif (/^func \(.*?\)\s+(\w+)\((.*?)\)\s+(.*?)({| )/) {
	$fn = $1;
	$fa = $2;
	$fr = $3;
	$fp = $f;
	$fat = "";
	$fant = "";
	foreach $z (split(",", $fa)) {
		@zz = split(" ", $z);
		$fat .= sprintf("- \x60%s\x60 - Type: \x60%s\x60 - @TODO\n", $zz[0], $zz[1]);
		$fant .= sprintf(" <%s>", $zz[0]);
	}
	$fp =~ s/%%ARGS%%/$fat/;
	$fp =~ s/%%FUNCARGS%%/$fant/;
	if ($fr =~ /State/) {
		$fr = 'state';
	}

	if ($fn =~ /^Helper/) {
		$fp =~ s/%%FUNC%%/\$$fr := $fn/g;
	} else {
		$fp =~ s/%%FUNC%%/\$state := $fn/g;
	}

	s/^(func)/$fp$1/;
}

