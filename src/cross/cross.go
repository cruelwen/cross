package main

import (
  "github.com/nsf/termbox-go"
  "math/rand"
  "time"
  "fmt"
  "os"
)

type typeMe struct {
  x,y int
  width,height int
}

func main() {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()
  termbox.SetInputMode(termbox.InputEsc)
  rand.Seed(time.Now().UnixNano())

  width, height := termbox.Size()
  me := typeMe {width / 2, height / 2, width, height}
  enemyList := make([]*typeEnemy,10)
  for i:=0;i<10;i++ {
    enemyList[i] = createEnemy(width, height)
    go enemyList[i].move()
  }

  go drawAll(&me,enemyList)
  for {
    me.moveWithKey()
  }
}

func drawAll(me *typeMe, enemyList []*typeEnemy) {
  for {
    const coldef = termbox.ColorDefault
    termbox.Clear(coldef, coldef)
    me.print()
    for _,enemy := range enemyList {
      enemy.print()
    }
    termbox.Flush()
    time.Sleep(100 * time.Millisecond)
  }
}

func (me typeMe) print() {
  const coldef = termbox.ColorDefault
  termbox.SetCell(me.x, me.y, '@', coldef, coldef)
}

func (me *typeMe) moveWithKey() {
  switch ev := termbox.PollEvent(); ev.Type {
  case termbox.EventKey:
    switch ev.Key {
    case termbox.KeyEsc:
      termbox.Close()
      fmt.Println("Bye")
      os.Exit(0)
    case termbox.KeyArrowLeft:
      me.x--
    case termbox.KeyArrowRight:
      me.x++
    case termbox.KeyArrowUp:
      me.y--
    case termbox.KeyArrowDown:
      me.y++
    default:
    }
  case termbox.EventError:
    panic(ev.Err)
  }
}

