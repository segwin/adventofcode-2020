package day13

import (
	"math/big"
	"strconv"
	"strings"
	"time"
)

type Bus struct {
	ID int
}

// DepartsIn returns the time until this bus's next departure after the given time,
// relative to time 0.
func (b Bus) DepartsIn(after time.Duration) time.Duration {
	frequency := b.ID
	afterMinutes := int((after).Minutes())

	nextDepartureMinutes := frequency - (afterMinutes % frequency)

	return time.Minute * time.Duration(nextDepartureMinutes)
}

type BusList struct {
	Buses []*Bus
}

func (l *BusList) Unmarshal(line string) error {
	for _, busStr := range strings.Split(line, ",") {
		if busStr == "x" {
			continue // this bus is out of service *or* represents a wildcard slot, if used in a Schedule
		}

		busID, err := strconv.Atoi(busStr)
		if err != nil {
			return err
		}

		l.Buses = append(l.Buses, &Bus{ID: busID})
	}

	return nil
}

type Schedule struct {
	Buses []*Subspace
}

func (s *Schedule) Unmarshal(line string) error {
	for i, busStr := range strings.Split(line, ",") {
		if busStr == "x" {
			continue // this bus is out of service *or* represents a wildcard slot, if used in a Schedule
		}

		busPeriod, err := strconv.Atoi(busStr)
		if err != nil {
			return err
		}

		s.Buses = append(s.Buses, NewSubspace(int64(busPeriod), int64(i)))
	}

	return nil
}

func (s *Schedule) FindLowestTime() (timeInMinutes int64, err error) {
	// don't bother checking length, just panic
	intersectionSubspace := s.Buses[0]

	for i := 1; i < len(s.Buses); i++ {
		intersectionSubspace, err = s.Buses[i].Intersect(intersectionSubspace)
		if err != nil {
			return 0, err
		}
	}

	return big.NewInt(0).Mod(big.NewInt(0).Neg(intersectionSubspace.Offset), intersectionSubspace.Coefficient).Int64(), nil
}
