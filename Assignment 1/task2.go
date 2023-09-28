package main

import "fmt"

type Observable interface {
	addObserver(observer Observer)
	removeObserver(observer Observer)
	notifyObservers()
	setAvailability(inStock bool)   //setAvailability  для того чтоб можно было устнаваливать состояния предмета(в наличии или нет)
}

type Observer interface {
	update(itemName string)
	getID() string
	display()
}

type Displayable interface {
	display() //Display позволяет наблюдателям отображать инфо о себе
}


type ItemSubject struct {
	observers   []Observer
	itemName    string
	isAvailable bool
}

func newItemSubject(itemName string) *ItemSubject {
	return &ItemSubject{
		itemName: itemName,
	}
}

func (i *ItemSubject) setAvailability(inStock bool) {
	i.isAvailable = inStock
	i.notifyObservers()
}

func (i *ItemSubject) addObserver(o Observer) {
	i.observers = append(i.observers, o)
}

func (i *ItemSubject) removeObserver(o Observer) {
	for idx, observer := range i.observers {
		if observer.getID() == o.getID() {
			// Remove the observer from the slice.
			i.observers = append(i.observers[:idx], i.observers[idx+1:]...)
			break
		}
	}
}

func (i *ItemSubject) notifyObservers() {
	for _, observer := range i.observers {
		observer.update(i.itemName)
	}
}

type ItemObserver struct {
	emailID string
}

func (c *ItemObserver) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.emailID, itemName)
}

func (c *ItemObserver) getID() string {
	return c.emailID
}

func (c *ItemObserver) display() {
	fmt.Printf("Observer: %s\n", c.emailID)
}

func main() {
	shirtItem := newItemSubject("Nike Shirt")

	observerFirst := &ItemObserver{emailID: "adammiguel2001@gmail.com"}
	observerSecond := &ItemObserver{emailID: "golangsuperpuper@gmail.com"}

	shirtItem.addObserver(observerFirst)
	shirtItem.addObserver(observerSecond)

	shirtItem.setAvailability(true) 

	fmt.Println("Observers:")
	for _, observer := range shirtItem.observers {
		observer.display()
	}
}
