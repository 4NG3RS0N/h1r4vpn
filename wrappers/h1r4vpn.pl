#!/usr/bin/env perl
use strict;
use warnings;
use FindBin;

exec "$FindBin::Bin/../h1r4vpn", @ARGV
    or die "Failed to execute h1r4vpn: $!\n";