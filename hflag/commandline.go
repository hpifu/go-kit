package hflag

import (
	"fmt"
	"net"
	"os"
	"time"
)

var CommandLine = NewFlagSet(os.Args[0])

func Lookup(name string) *Flag {
	return CommandLine.Lookup(name)
}

func Set(name string, val string) error {
	return CommandLine.Set(name, val)
}

func Visit(callback func(f *Flag)) {
	CommandLine.Visit(callback)
}

func VisitAll(callback func(f *Flag)) {
	CommandLine.VisitAll(callback)
}

func Parsed() bool {
	return CommandLine.Parsed()
}

func NArg() int {
	return CommandLine.NArg()
}

func Args() []string {
	return CommandLine.Args()
}

func Arg(i int) string {
	return CommandLine.Arg(i)
}

func NFlag() int {
	return CommandLine.NFlag()
}

func PrintDefaults() {
	CommandLine.PrintDefaults()
}

func Bool(name string, defaultValue bool, usage string) *bool {
	return CommandLine.Bool(name, defaultValue, usage)
}

func Int(name string, defaultValue int, usage string) *int {
	return CommandLine.Int(name, defaultValue, usage)
}

func Int64(name string, defaultValue int64, usage string) *int64 {
	return CommandLine.Int64(name, defaultValue, usage)
}

func Uint(name string, defaultValue uint, usage string) *uint {
	return CommandLine.Uint(name, defaultValue, usage)
}

func Uint64(name string, defaultValue uint64, usage string) *uint64 {
	return CommandLine.Uint64(name, defaultValue, usage)
}

func String(name string, defaultValue string, usage string) *string {
	return CommandLine.String(name, defaultValue, usage)
}

func Duration(name string, defaultValue time.Duration, usage string) *time.Duration {
	return CommandLine.Duration(name, defaultValue, usage)
}

func Float64(name string, defaultValue float64, usage string) *float64 {
	return CommandLine.Float64(name, defaultValue, usage)
}

func IntSlice(name string, defaultValue []int, usage string) *[]int {
	return CommandLine.IntSlice(name, defaultValue, usage)
}

func StringSlice(name string, defaultValue []string, usage string) *[]string {
	return CommandLine.StringSlice(name, defaultValue, usage)
}

func Time(name string, defaultValue time.Time, usage string) *time.Time {
	return CommandLine.Time(name, defaultValue, usage)
}

func IP(name string, defaultValue net.IP, usage string) *net.IP {
	return CommandLine.IP(name, defaultValue, usage)
}

func BoolVar(v *bool, name string, defaultValue bool, usage string) {
	CommandLine.BoolVar(v, name, defaultValue, usage)
}

func IntVar(v *int, name string, defaultValue int, usage string) {
	CommandLine.IntVar(v, name, defaultValue, usage)
}

func Int64Var(v *int64, name string, defaultValue int64, usage string) {
	CommandLine.Int64Var(v, name, defaultValue, usage)
}

func UintVar(v *uint, name string, defaultValue uint, usage string) {
	CommandLine.UintVar(v, name, defaultValue, usage)
}

func Uint64Var(v *uint64, name string, defaultValue uint64, usage string) {
	CommandLine.Uint64Var(v, name, defaultValue, usage)
}

func StringVar(v *string, name string, defaultValue string, usage string) {
	CommandLine.StringVar(v, name, defaultValue, usage)
}

func DurationVar(v *time.Duration, name string, defaultValue time.Duration, usage string) {
	CommandLine.DurationVar(v, name, defaultValue, usage)
}

func Float64Var(v *float64, name string, defaultValue float64, usage string) {
	CommandLine.Float64Var(v, name, defaultValue, usage)
}

func IntSliceVar(v *[]int, name string, defaultValue []int, usage string) {
	CommandLine.IntSliceVar(v, name, defaultValue, usage)
}

func StringSliceVar(v *[]string, name string, defaultValue []string, usage string) {
	CommandLine.StringSliceVar(v, name, defaultValue, usage)
}

func TimeVar(v *time.Time, name string, defaultValue time.Time, usage string) {
	CommandLine.TimeVar(v, name, defaultValue, usage)
}

func IPVar(v *net.IP, name string, defaultValue net.IP, usage string) {
	CommandLine.IPVar(v, name, defaultValue, usage)
}

func GetInt(name string) int {
	return CommandLine.GetInt(name)
}

func GetFloat(name string) float64 {
	return CommandLine.GetFloat(name)
}

func GetString(name string) string {
	return CommandLine.GetString(name)
}

func GetDuration(name string) time.Duration {
	return CommandLine.GetDuration(name)
}

func GetBool(name string) bool {
	return CommandLine.GetBool(name)
}

func GetIntSlice(name string) []int {
	return CommandLine.GetIntSlice(name)
}

func GetStringSlice(name string) []string {
	return CommandLine.GetStringSlice(name)
}

func GetIP(name string) net.IP {
	return CommandLine.GetIP(name)
}

func GetTime(name string) time.Time {
	return CommandLine.GetTime(name)
}

func Parse() error {
	if CommandLine.Lookup("help") == nil && CommandLine.Lookup("h") == nil {
		CommandLine.AddFlag("help", "show usage", Shorthand("h"), Type("bool"))
	}
	if err := CommandLine.Parse(os.Args[1:]); err != nil {
		if CommandLine.GetBool("help") {
			fmt.Println(CommandLine.Usage())
			os.Exit(0)
		}
		return err
	}
	if CommandLine.GetBool("help") {
		fmt.Println(CommandLine.Usage())
		os.Exit(0)
	}

	return nil
}

func AddFlag(name string, usage string, opts ...FlagOption) {
	CommandLine.AddFlag(name, usage, opts...)
}

func AddPosFlag(name string, usage string, opts ...FlagOption) {
	CommandLine.AddPosFlag(name, usage, opts...)
}

func Usage() string {
	return CommandLine.Usage()
}
