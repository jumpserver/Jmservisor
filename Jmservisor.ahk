if(A_Args[1]="chrome"){
    Run,%2% --app=%3% --start-maximized
    WinActivate, ahk_exe chrome.exe
    BlockInput,on
    Sleep, 10000
    SetKeyDelay, 100
    Send, {Tab}
    Sleep, 300
    Send, %4% ;用户名
    Sleep, 300
    Send, {Tab}
    Sleep, 300
    SendRaw, %5% ;密码
    Sleep, 300
    Send, {Tab}
    Sleep, 300
    Send, {Enter}
    BlockInput,off
    return
}

if(A_Args[1]="mysql_workbench"){
    Run, %2%
    WinActivate, ahk_exe MySQLWorkbench.exe
    BlockInput,on
    Sleep, 10000
    SetKeyDelay, 100
    Send , ^u
    Send, {Tab} ;跳过保存连接选项
    Send, {Tab} ;跳过连接选型选项
    Send, {Tab} ;跳过Tag选项
    Sleep, 300
    Send, ^a
    Send, {Delete}
    Send, %3% ;IP
    Sleep, 300
    Send, {Tab}
    Send, ^a
    Send, {Delete}
    Send, %4% ;端口
    Sleep, 300
    Send, {Tab}
    Send, ^a
    Send, {Delete}
    Send, %6% ;用户名
    Sleep, 300
    Send, {Tab}
    Send, {Enter}
    Sleep, 300
    SendRaw, %7% ;密码
    Sleep, 300
    Send, {Enter}
    Sleep, 300
    Send, {Tab}
    Send, {Tab}
    Send, ^a
    Send, {Delete}
    Sleep, 300
    Send, %5% ;默认数据库
    Sleep, 300
    Send, {Enter}
    BlockInput,off
    return
}

if(A_Args[1]="vmware_client"){
    Run,%2%
    WinActivate, ahk_exe VpxClient.exe
    BlockInput,on
    Sleep, 10000
    SetKeyDelay, 100
    ; 第二次登录 有IP记录情况
    Send, {Tab}
    Send, {Tab}
    Send, {Tab}
    Send, {Tab}
    Send, {Tab}
    Send, %3% ;IP
    Send, {Tab}
    Sleep, 300
    Send, %4% ;用户名
    Send, {Tab}
    Sleep, 300
    SendRaw, %5% ;密码
    Send, {Enter}
    BlockInput,off
    return
}

if(A_Args[1]="custom"){
	  if(!FileExist(A_Args[2]) or InStr(FileExist(A_Args[2]),"D")){
		    return
  	}
	  params := ""
	  if(A_Args.length() > 2){
		    for index,param in A_Args{
			      if(index > 2){
				        params := params " " param
			      }
		    }
	  }
	  Run,%2% %params%
	  return
}
return
