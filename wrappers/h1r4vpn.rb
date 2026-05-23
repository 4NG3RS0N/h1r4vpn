#!/usr/bin/env ruby

require 'pathname'

bin = Pathname.new(__FILE__).realpath.dirname + '../h1r4vpn'

exec(bin.to_s, *ARGV)