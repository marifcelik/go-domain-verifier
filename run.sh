runCommand="go run ."

if [ "$1" != "-n"]; then
    runCommand="watchexec -rc -e go -- $runCommand"
fi

$runCommand