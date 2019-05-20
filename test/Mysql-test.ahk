Run,%1%
WinActivate, ahk_exe MySQLWorkbench.exe

BlockInput,on
Sleep, 10000
SetKeyDelay, 30
Send , ^u
Send, {Tab} ;跳过保存连接选项
Send, {Tab} ;跳过连接选型选项
Send, {Tab} ;跳过Tag选项

sleep, 100
Send, ^a ;输入IP
Send, %2%

sleep, 100
Send, {Tab};
Send, ^a ;输入端口
Send, %3%

sleep, 100
Send, {Tab};
Send, ^a ;输入用户名
Send, %4%

sleep, 100
Send, {Tab};
Send, {Enter} ;输入密码
Send, %5%
Send, {Enter}

sleep, 100
Send, {Tab} 
Send, {Tab} 
Send, {Tab} 
Send, {Tab} 
Send, {Enter} ;确定
BlockInput,off