package main

import (
  "github.com/nsf/termbox-go"
  "math/rand"
  "flag"
  "time"
  "fmt"
  "strconv"
  "os"
)

type typeMe struct {
  x,y int
  width,height int
}

func main() {
  var startupNum = flag.Int("n", 0, "startup enemies num")
  flag.Parse()
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  defer termbox.Close()
  termbox.SetInputMode(termbox.InputEsc)
  rand.Seed(time.Now().UnixNano())

  width, height := termbox.Size()
  me := typeMe {width / 2, height / 2, width, height}
  enemyList := [1000]*typeEnemy {}
  enemyNum := 0
  for ;enemyNum<*startupNum;enemyNum++ {
    enemyList[enemyNum] = createEnemy(width, height)
    go enemyList[enemyNum].move()
  }

  go drawAll(&me,&enemyList,&enemyNum)
  go enemyMaster(&enemyList,&enemyNum)
  for {
    me.moveWithKey()
  }
}

func enemyMaster(enemyList *[1000]*typeEnemy,enemyNum *int) {
  width, height := termbox.Size()
  for ;*enemyNum < 1000;*enemyNum++ {
    enemyList[*enemyNum] = createEnemy(width, height)
    go enemyList[*enemyNum].move()
    time.Sleep(100 * time.Millisecond)
  }
}

func drawAll(me *typeMe, enemyList *[1000]*typeEnemy,enemyNum *int) {
  for {
    const coldef = termbox.ColorDefault
    termbox.Clear(coldef, coldef)
    // player
    me.print()
    // score
    score := strconv.Itoa(*enemyNum)
    for p, c := range score {
      termbox.SetCell(p,0,c,coldef,coldef)
    }

    // enemy
    for i:=0;i<*enemyNum;i++ {
      enemyList[i].print()
      if me.x == enemyList[i].x && me.y == enemyList[i].y {
        termbox.Close()
        fmt.Println("Score:",*enemyNum)
        os.Exit(0)
      }
    }
    termbox.Flush()
    time.Sleep(30 * time.Millisecond)
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

