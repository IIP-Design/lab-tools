---
title: June 25, 2021 - Security Incident Postmortem Template
subtitle: A More Detailed Description of the Issue
toc: true # Enables the table of contents, should be set to facilitate generating PDF reports
date: 2021-07-01 # The date of the final official report
---

## Incident Summary

Write a summary of the incident. If they only read this section, the reader should come away with a basic understanding of the what happened and how it was fixed. Make sure to include:

- a general description of what happened
- an brief assessment of why it happened
- the severity of the incident (both actual and theoretical)
- an outline of remedial actions taken
- how long the impact lasted (i.e. when we were first notified, when did we considered it resolved)

## Lead-Up

Describe the sequence of events that led to the incident. Provide the reader with the necessary background to understand the conditions that allowed the incident to occur. Include relevant changes to the system or interactions that contributed to causing the incident.

Beyond providing a play-by-play of changes to the system, this section should communicate the team's intent, explain its decision making process, and highlight necessary tradeoffs that led to the incident.

## Fault

Winnow down the lead-up to a single point of failure. Think of this as the straw that broke the camel's back. Describe how the change that was implemented didn't work as expected.

If possible, attach screenshots, code snippets, or relevant data visualizations that illustrate the fault.

## Impact

Describe how the incident impacted internal and external users during the incident. Identify as precisely as possible how many users may have been adversely impacted and for how long. Include how many support cases were raised or complaints received.

If the incident in question is a security vulnerability responsible disclosure, describe the potential impact had the vulnerability been exploited. Furthermore, describe any efforts to conduct forensics and determine unreported/unnoticed impacts to the system.

## Detection

Describe how the team learned of the issue:

- When did the team detect the incident?
- How did they know it was happening?
- Did detection originate internal or was it reported by users?
- How could we improve time-to-detection?

## Response

Describe how the team decided on a course of action and what immediate steps were taken to address the issue.

- Who responded to the incident?
- When did they respond and what did they do?
- Note any delays or obstacles to responding.

## Recovery

Describe how the service was restored and when the incident was deemed over.

Depending on the scenario, consider these questions:
  - How could you improve time to mitigation?
  - How could you have cut that time by half?

## Timeline

Detail the incident timeline. Include any notable lead-up events, any starts of activity, the first known impact, and escalations. Note any decisions points and when the incident ended. 

Times should be listed using 24 hour notation and based on the EDT/EST timezone. If the incident spans multiple days, begin each day with a date heading in the format: month spelled out, two digit day, four digit year.

**EXAMPLE:**

**June 25, 2021:**

11:48 - K8S 1.9 upgrade of control plane is finished 

12:46 - Upgrade to v1.9 completed, including cluster-auto scaler and the BuildEng scheduler instance 

## Root Cause Identification

Attempt to identify the root cause of the issue using the "five whys" technique. Namely,

- Begin with a description of the impact and ask why it occurred. 
- Note the impact that it had.  
- Ask why this happened, and why it had the resulting impact. 
- Then, continue asking “why” until you arrive at a root cause.
- List the "whys" in your postmortem documentation.

## Root Cause

Note the final root cause of the incident. Specifically, this is the thing that needs to change in order to prevent this class of incident from happening again. This should be encapsulated in a very concise one to two sentences.

## Backlog Check

A clear-eyed assessment of the backlog can shed light on past decisions around priority and risk. Review your engineering backlog to find out if there was any planned or unplanned work there that could have prevented this incident, or at least reduced its impact.

## Recurrence

With the root cause identified, look back at previous incidents. Could they have the same root cause?

If yes, note what mitigation was attempted in those incidents and ask why this incident occurred again.

## Corrective Actions

Outline the steps taken to address the issue. Include immediate action to patch security vulnerabilities/restore service.

Also, describe the corrective action ordered to prevent or mitigate this class of incident in the future. Note who is responsible, what is the deadline to complete the work, and where that work is being tracked.