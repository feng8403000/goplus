package main

import spx "github.com/goplus/spx"

type index struct {
	spx.Game
	butTen butTen
	times  int
}
type butTen struct {
	spx.Sprite
	*index
}

func (this *index) MainEntry() {
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\index.gmx:6
	spx.Gopt_Game_Run(this, "res", &spx.Config{Title: "10秒真男人游戏"})
}
func main() {
	spx.Gopt_Game_Main(new(index))
}
func (this *butTen) Main() {
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:1
	this.OnTouched__1(func() {
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:2
		this.Destroy()
	})
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:4
	this.OnStart(func() {
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:5
		for {
			spx.Sched()
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:6
			for
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:6
			i := 0; i < 11;
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:6
			i += 1 {
				spx.Sched()
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:7
				this.times = i
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:8
				this.Wait(0.01)
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:9
				spx.Gopt_Sprite_Clone__0(this)
			}
		}
	})
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:13
	this.OnClick(func() {
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:14
		if this.times == 10 {
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:15
			this.Say("真男人", 2)
		} else {
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:17
			this.Say(this.times, 2.0)
		}
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:19
		this.times = 0
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\butTen.spx:20
		spx.Gopt_Sprite_Clone__0(this)
	})
}

type stop struct {
	spx.Sprite
	*index
}

func (this *stop) Main() {
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\stop.spx:1
	this.OnKey__0(spx.KeySpace, func() {
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\stop.spx:2
		this.Say(this.times)
	})
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\stop.spx:4
	this.OnClick(func() {
//line E:\goplus\gopGame\spx\tutorial\07_Ten_test\stop.spx:5
		this.Say(this.times)
	})
}
