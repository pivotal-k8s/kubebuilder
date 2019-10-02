package controllers

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	batch "tutorial.kubebuilder.io/project/api/v1"
)

var _ = Describe("CronjobController", func() {
	Describe("getNextSchedule", func() {
		It("accurately calculates lastMissed", func() {

			cronJob := batch.CronJob{
				ObjectMeta: metav1.ObjectMeta{}
				Spec: batch.CronJobSpec{
					Schedule: "* * * * *",
				},
				Status: batch.CronJobStatus{},
			}

			_, _, err := getNextSchedule(&cronJob, time.Now())
			Expect(err).NotTo(HaveOccurred())
			// Expect(lastMissed).To(BeNil())
			// Expect(next).To(BeNil())
		})
	})

})
