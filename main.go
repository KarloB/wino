package main

import "fmt"
import "sort"

type containerType int

const coef = 0.9

const (
	inox containerType = iota + 1
	wood
)

type wine int

const (
	skrlet wine = iota + 1
	grasevina
	riesling
	pinot
	sedminac
)

type Wine struct {
	Type     wine
	Quantity int
}

type Container struct {
	ID           int
	Type         containerType
	Capacity     int
	RealCapacity int
	UsedCapacity int
	PreferedWine wine
}

var containers = []Container{
	Container{ID: 1, Type: inox, Capacity: 2000},
	Container{ID: 2, Type: inox, Capacity: 1100},
	Container{ID: 3, Type: inox, Capacity: 1100},
	Container{ID: 4, Type: inox, Capacity: 1000},
	Container{ID: 5, Type: inox, Capacity: 1000},
	Container{ID: 6, Type: inox, Capacity: 1000},
	Container{ID: 7, Type: inox, Capacity: 1000},
	Container{ID: 8, Type: inox, Capacity: 1000},
	Container{ID: 9, Type: wood, Capacity: 1600, PreferedWine: grasevina},
	Container{ID: 10, Type: wood, Capacity: 550, PreferedWine: grasevina},
	Container{ID: 11, Type: wood, Capacity: 520, PreferedWine: skrlet},
	Container{ID: 12, Type: wood, Capacity: 440, PreferedWine: pinot},
	Container{ID: 13, Type: wood, Capacity: 430, PreferedWine: pinot},
	Container{ID: 14, Type: inox, Capacity: 620},
	Container{ID: 15, Type: inox, Capacity: 400},
}

var wines = []Wine{
	Wine{Type: skrlet, Quantity: 1500},
	Wine{Type: grasevina, Quantity: 3000},
	Wine{Type: riesling, Quantity: 1000},
	Wine{Type: pinot, Quantity: 3000},
	Wine{Type: sedminac, Quantity: 2000},
}

type Result struct {
	Wine              Wine
	Containers        []Container
	ContainerCapacity int
	UnusedCapacity    int
}

func main() {
	var result []Result
	usedContainers := make(map[int]bool)
	var totalwine int
	var totalContainer int

	for i, c := range containers {
		usedCapacity := float64(c.Capacity) * coef
		containers[i].RealCapacity = int(usedCapacity)
		totalContainer += int(usedCapacity)
	}
	sort.Sort(SortContainerBySmallest(containers))

	for _, wine := range wines[0:1] {
		remaining := wine.Quantity
		var cont []Container
		var preferedContainer []Container
		// var containerCapacity int

		for _, c := range containers {
			if c.PreferedWine == wine.Type {
				if ok, _ := usedContainers[c.ID]; ok {
					continue
				}
				preferedContainer = append(preferedContainer, c)
			}
		}

		// for {

		// 	container := findPreferedContainer

		// 	if remaining <= 0 {
		// 		break
		// 	}
		// }

		for _, c := range containers {
			if ok, _ := usedContainers[c.ID]; ok {
				continue
			}
			if c.PreferedWine != 0 && c.PreferedWine != wine.Type {
				continue
			}
			remaining, cont = addWineToContainer(c, remaining, cont)
			usedContainers[c.ID] = true
			if remaining <= 0 {
				break
			}
		}

		var containerCapacity int
		for _, c := range cont {
			containerCapacity += c.Capacity
		}

		entry := Result{
			Wine:              wine,
			Containers:        cont,
			ContainerCapacity: containerCapacity,
			UnusedCapacity:    0,
		}
		result = append(result, entry)
	}

	fmt.Println("Total wine: ", totalwine)
	fmt.Println("Container capacity: ", totalContainer)

	for _, r := range result {
		fmt.Printf("Wine: %+v\n", r.Wine)
		fmt.Printf("Containers capacity: %d\n", r.ContainerCapacity)
		for _, c := range r.Containers {
			fmt.Printf("Container: %+v\n", c)
		}
		fmt.Println("- - - - - - - - - -")
	}
}

func addWineToContainer(container Container, remaining int, containers []Container) (int, []Container) {
	var usedCapacity int
	// containerCapacity += container.Capacity
	remaining = remaining - container.RealCapacity
	if remaining > 0 {
		usedCapacity = container.RealCapacity
	} else {
		usedCapacity = container.RealCapacity + remaining
	}
	container.UsedCapacity = usedCapacity
	containers = append(containers, container)
	return remaining, containers
}

func findClosestContainer(remaining int, containers []Container) Container {
	var result Container

	return result
}

// var preferedContainer Container
// for _, c := range containers {
// 	if c.PreferedWine == wine.Type {
// 		preferedContainer = c
// 		break
// 	}
// }
// if preferedContainer.ID != 0 {

// }

type SortContainerBySmallest []Container

func (a SortContainerBySmallest) Len() int           { return len(a) }
func (a SortContainerBySmallest) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortContainerBySmallest) Less(i, j int) bool { return a[i].Capacity < a[j].Capacity }
