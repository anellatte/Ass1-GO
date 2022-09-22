package main

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

func (j *JobSite) subscribe(observer Observer) {
	j.subscribers = append(j.subscribers, observer)
}

func (j *JobSite) unsubscribe(observer Observer) {
	for m, n := range j.subscribers {
		if n == observer {
			j.subscribers = append(j.subscribers[:m], j.subscribers[m+1:]...)
		}
	}
}

func (j JobSite) sendAll() {
	for _, subscriber := range j.subscribers {
		subscriber.handleEvent(j.vacancies)
	}
}

func (j *JobSite) addVacancy(vacancy string) {
	j.vacancies = append(j.vacancies, vacancy)
	j.sendAll()
}

func (j *JobSite) removeVacancy(vacancy string) {
	for m, n := range j.vacancies {
		if n == vacancy {
			j.vacancies = append(j.vacancies[:m], j.vacancies[m+1:]...)
		}
	}

	j.sendAll()
}
