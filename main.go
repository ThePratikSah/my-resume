package main

import (
	"log"
	"os"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	cfg := config.NewBuilder().
		WithPageSize("A4").
		WithLeftMargin(15).
		WithTopMargin(15).
		WithRightMargin(15).
		WithBottomMargin(15).
		Build()

	m := maroto.New(cfg)

	buildHeader(m)
	buildSummary(m)
	buildSkills(m)
	buildExperience(m)
	buildEducation(m)

	doc, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("resume.pdf", doc.GetBytes(), 0644); err != nil {
		log.Fatal(err)
	}
}

func buildHeader(m core.Maroto) {
	m.AddRow(10,
		text.NewCol(12, "Pratik Sah", props.Text{
			Style: fontstyle.Bold,
			Size:  22,
			Align: "center",
			Family: fontfamily.Arial,
		}),
	)

	m.AddRow(8,
		text.NewCol(12, "Sr. Associate (Backend Engineer)  • Node.js  •  React  • Distributed Systems", props.Text{
			Size:  11,
			Align: "center",
			Family: fontfamily.Arial,
		}),
	)

	m.AddRow(6,
		text.NewCol(12, "Bengaluru, India  |  pratiksah@hotmail.com  |  (+91)-870-910-5800", props.Text{
			Size:  9,
			Align: "center",
			Family: fontfamily.Arial,
		}),
	)

	m.AddRow(6,
		text.NewCol(12, "linkedin.com/in/pratiksah  |  github.com/ThePratikSah", props.Text{
			Size:  9,
			Align: "center",
			Family: fontfamily.Arial,
		}),
	)

	m.AddRow(3)
}

func buildSummary(m core.Maroto) {
	m.AddRow(8,
		text.NewCol(12, "SUMMARY", props.Text{
			Style: fontstyle.Bold,
			Size:  13,
			Family: fontfamily.Arial,
		}),
	)
	m.AddRow(2)
	m.AddRow(30,
		text.NewCol(12, "Backend engineer with 5+ years building high-throughput distributed systems using Node.js, Go, and AWS, with prior experience in full-stack (React, React Native). Scaled Razorpay billing from 100K to 1.5M+ daily transactions using async architectures. Currently working with Emirates NBD (via Synechron) on banking-grade microservices.", props.Text{
			Size:  10,
			Family: fontfamily.Arial,
		}),
	)
}

func buildSkills(m core.Maroto) {
	m.AddRow(2)
	m.AddRow(8,
		text.NewCol(12, "SKILLS", props.Text{
			Style: fontstyle.Bold,
			Size:  13,
			Family: fontfamily.Arial,
		}),
	)
	m.AddRow(2)

	skills := []struct{ label, items string }{
		{"Languages:", "Go, Node.js, TypeScript, Python"},
		{"Frameworks:", "NestJS, Express"},
		{"Frontend:", "React, React Native"},
		{"Cloud & Containers:", "AWS, GCP, Docker, AWS CDK"},
		{"Messaging:", "Kafka, RabbitMQ, SNS/SQS"},
		{"Databases:", "MongoDB, MySQL, DynamoDB, Redis"},
		{"Observability:", "ELK Stack"},
		{"Architecture:", "Microservices, Event-Driven, Serverless"},
	}

	for _, s := range skills {
		m.AddRow(5,
			col.New(3).Add(
				text.New(s.label+" ", props.Text{
					Style: fontstyle.Bold,
					Size:  10,
					Align: "right",
					Family: fontfamily.Arial,
				}),
			),
			col.New(9).Add(
				text.New(s.items, props.Text{
					Size:  10,
					Family: fontfamily.Arial,
				}),
			),
		)
	}
}

func buildExperience(m core.Maroto) {
	m.AddRow(3)
	m.AddRow(8,
		text.NewCol(12, "EXPERIENCE", props.Text{
			Style: fontstyle.Bold,
			Size:  13,
			Family: fontfamily.Arial,
		}),
	)

	experiences := []struct {
		company, role, period, details string
	}{
		{
			"Synechron Technologies",
			"Sr. Associate (Backend Engineer)",
			"Nov 2025 — Present  |  Bengaluru, India (Client: Emirates NBD, Dubai)",
			"Building Node.js microservices deployed on OpenShift (OCP) + Docker for banking workloads.\nDesigned ELK-based observability pipelines for centralized logging, monitoring, and alerting.\nIntegrated Kafka for event-driven communication, improving system resilience and decoupling.\nCollaborating with ENBD teams on CI/CD, container lifecycle, and platform standards.",
		},
		{
			"Razorpay",
			"Product Development Engineer II",
			"Jan 2024 — Nov 2025  |  Bengaluru",
			"Led Node.js to Go migration, improving throughput by 45% and cutting latency by 30%.\nBuilt distributed async pipelines (SNS/SQS/S3) handling 1.5M+ events/day.\nExecuted zero-downtime MongoDB to MySQL migration.\nMentored engineers and drove design reviews across microservices.",
		},
		{
			"Ethereal Covenant",
			"Software Development Engineer",
			"May 2023 — Dec 2023",
			"Developed full-stack apps using React Native, Node.js, MongoDB, and AWS.\nContributed to UI components and frontend integration alongside backend services.\nImproved response time ~25% via caching and query optimization.",
		},
		{
			"Altsys",
			"Senior Software Engineer",
			"Aug 2022 — Sept 2023  |  Bengaluru",
			"Owned backend services using Node.js, PostgreSQL, MongoDB and supported frontend integrations.\nIntroduced caching + CI/CD pipelines, improving reliability and latency.",
		},
		{
			"Deqode",
			"Solutions Engineer",
			"Nov 2021 — Aug 2022  |  Pune",
			"Built microservices using Node.js, RabbitMQ, MongoDB.\nMigrated legacy systems to event-driven architecture (NestJS).",
		},
		{
			"CodeBucket Solutions",
			"Full Stack Developer",
			"Dec 2020 — Nov 2021  |  Kochi",
			"Built full-stack applications using React, Node.js, MySQL, NGINX.\nDeveloped real-time systems using Socket.io, Firebase, and Redis.\nDelivered production-ready UI and backend systems.",
		},
	}

	for _, e := range experiences {
		m.AddRow(2)
		m.AddRow(7,
			text.NewCol(12, e.company+" — "+e.role, props.Text{
				Style: fontstyle.Bold,
				Size:  11,
				Family: fontfamily.Arial,
			}),
		)
		m.AddRow(5,
			text.NewCol(12, e.period, props.Text{
				Size:  9,
				Family: fontfamily.Arial,
			}),
		)
		m.AddRow(2)
		m.AddRow(16,
			text.NewCol(12, e.details, props.Text{
				Size:  10,
				Family: fontfamily.Arial,
			}),
		)
	}
}

func buildEducation(m core.Maroto) {
	m.AddRow(3)
	m.AddRow(8,
		text.NewCol(12, "EDUCATION", props.Text{
			Style: fontstyle.Bold,
			Size:  13,
			Family: fontfamily.Arial,
		}),
	)
	m.AddRow(2)
	m.AddRow(6,
		text.NewCol(12, "B.Tech, Computer Science", props.Text{
			Size:  10,
			Family: fontfamily.Arial,
			Style: fontstyle.Bold,
		}),
	)
	m.AddRow(5,
		text.NewCol(12, "Aryabhatta Knowledge University (2016 - 2020)", props.Text{
			Size:  10,
			Family: fontfamily.Arial,
		}),
	)
}
