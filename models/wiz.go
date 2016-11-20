package models

type Wiz struct {
	Id int
	Content string
	CreatedAt int
	ZombieId int
	Zombie *Zombie
}
