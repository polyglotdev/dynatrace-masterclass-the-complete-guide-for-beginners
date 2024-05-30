# Dynatrace Masterclass - The Complete Guide for Beginners

[course link](https://www.udemy.com/course/dynatrace-learning-tutorial)

<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/TYPESCRIPT-black?style=for-the-badge&logo=TypeScript&logoColor=3178C6" alt="TYPESCRIPT logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/GIT-black?style=for-the-badge&logo=GIT&logoColor=F05032" alt="GIT logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/JAVASCRIPT-black?style=for-the-badge&logo=JavaScript&logoColor=F7DF1E" alt="JAVASCRIPT logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/SH SCRIPTS-black?style=for-the-badge&logo=GNU Bash&logoColor=white" alt="SH logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/RUST-black?style=for-the-badge&logo=RUST&logoColor=white" alt="RUST logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/PYTHON-black?style=for-the-badge&logo=python&logoColor=gold" alt="PYTHON logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/Dart-black?style=for-the-badge&logo=Dart&logoColor=0175C2" alt="Dart logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/Kotlin-black?style=for-the-badge&logo=Kotlin&logoColor=7F52FF" alt="Kotlin logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/HTML-black?style=for-the-badge&logo=HTML5&logoColor=E34F26" alt="HTML logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/C Sharp-black?style=for-the-badge&logo=C sharp&logoColor=512BD4" alt="C logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/GO-black?style=for-the-badge&logo=Go&logoColor=00ADD8" alt="GO logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/JAVA-black?style=for-the-badge&logo=openjdk&logoColor=F37626" alt="JAVA logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/DOCKER-black?style=for-the-badge&logo=Docker&logoColor=2496ED" alt="DOCKER logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/PHP-black?style=for-the-badge&logo=PHP&logoColor=777BB4" alt="PHP logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/SWIFT-black?style=for-the-badge&logo=swift&logoColor=F05138" alt="SWIFT logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/CSS3-black?style=for-the-badge&logo=CSS3&logoColor=1572B6" alt="CSS3 logo"/>
<!-- trunk-ignore(markdownlint/MD033) -->
<img src="https://img.shields.io/badge/react-black?style=for-the-badge&logo=react&logoColor=61DAFB" alt="react logo"/>

[Dynatrace Tenant](https://asw86539.live.dynatrace.com/ui/dashboards?gtf=-2h&gf=all)

## Section 2: The Basics

- Create an account with Dynatrace
- Logging to Dynatrace
- Important: Latest interface vs classic View
- Deploy Dynatrace OneAgent

## Section 3: Infrastructure Observability

### Host Performance Data

#### Host Problem Alerting

Dynatrace can alert you when a host is having problems. You can configure the alerting rules in the `Settings` -> `Anomaly Detection` -> `Hosts` -> `Host availability` section.

![Host Problems](./images/host-errors.png)

[Resource Events](https://docs.dynatrace.com/docs/platform/davis-ai/basics/events/event-types/resource-events)

#### Host Availability

Host availability is the percentage of time that a host is available. It is calculated as the time the host is available divided by the total time.

- When its running its blue and when its offline its red.
- Offline does not mean the host is down, it could be that the agent is not running or there could be a communication error.
- It will also show the Maintenance window if one is set.
  - So if you know the host is going to be down for maintenance you can set a maintenance window so that you don't get alerts.

### Application Monitoring

[Docs](https://docs.dynatrace.com/docs/manage/subscriptions-and-licensing/monitoring-consumption-classic/application-and-infrastructure-monitoring)

We have an option to monitor the application and infrastructure. The application monitoring is done by the OneAgent. The OneAgent is a lightweight agent that is installed on the host and it collects data about the host and the applications running on the host. It has two modes:

- Full-stack monitoring: This is the default mode and it monitors the host and the applications running on the host.
- Infrastructure monitoring: This mode only monitors the host and not the applications running on the host.

The Full-stack option is going to be more expensive but has added benefits. The infrastructure monitoring is going to be cheaper but you will not get the application monitoring. It can be 3x less expensive to use the infrastructure monitoring.

## Section 4: Cloud Automation
