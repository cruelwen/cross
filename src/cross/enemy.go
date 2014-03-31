package main

import (
  "github.com/nsf/termbox-go"
  "math/rand"
  "time"
)

type typeEnemy struct {
  x,y int
  dx,dy int
  width,height int
  sleepTime time.Duration
  isEnd bool
}

func createEnemy(width,height int) (enemy *typeEnemy) {
  enemy = new(typeEnemy)
  enemy.width = width
  enemy.height = height
  enemy.isEnd = false
  xOrY := rand.Intn(2)
  if xOrY == 0 {
    enemy.dx = 0
    enemy.dy = rand.Intn(2) * 2 - 1
    enemy.x = rand.Intn(width)
    if enemy.dy == 1 {
      enemy.y = 0
    } else {
      enemy.y = height
    }
  } else {
    enemy.dy = 0
    enemy.dx = rand.Intn(2) * 2 - 1
    enemy.y = rand.Intn(height)
    if enemy.dx == 1 {
      enemy.x = 0
    } else {
      enemy.x = width
    }
  }
  enemy.sleepTime = time.Duration(rand.Intn(5) * 100 + 100)
  return
}

func (enemy *typeEnemy) next() (isEnd bool) {
  enemy.x += enemy.dx
  enemy.y += enemy.dy
  if enemy.x > enemy.width || enemy.x < 0 || enemy.y > enemy.height || enemy.y < 0 {
    isEnd = true
    enemy.isEnd = true
  }
  return
}

func (enemy typeEnemy) sleep() {
  time.Sleep(enemy.sleepTime * time.Millisecond)
}

func (enemy typeEnemy) print() {
  const coldef = termbox.ColorDefault
  termbox.SetCell(enemy.x, enemy.y, 'O', coldef, coldef)
}

func (enemy *typeEnemy) move() {
  for {
    enemy.next()
    enemy.sleep()
  }
}
