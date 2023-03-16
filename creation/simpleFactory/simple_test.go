package simpleFactory

import "testing"

//简单工厂设计模式测试

func TestType1(t *testing.T){
	api:=NewAPI(1)
	s:=api.Say("Sttch")
	if s!="Hi,Sttch"{
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T){
	api:=NewAPI(2)
	s:=api.Say("Sttch")
	if s!="Hello,Sttch"{
		t.Fatal("type2 test fail")
	}
}
