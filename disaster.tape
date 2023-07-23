Output static/demo.gif

Require disaster

Set Shell "bash"

Set WindowBar Colorful
Set BorderRadius 10
Set FontSize 12
Set Width 1200
Set Height 600
Set Margin 60

Show
Type@100ms "disaster"
Enter
Sleep 2s
Type "disaster categories --output table"
Enter
Sleep 3s
Type "disaster events --output table"
Enter
Sleep 6s
Type "clear"
Enter
Type "disaster events --output json | jq"
Enter
Sleep 6s
