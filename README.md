# Loyal Customer Tracker

This project identifies loyal customers from website access logs based on two criteria:
1. They visit the website on both days
2. They view at least 2 different pages across those days

## Problem Statement

The system logs user website visits with: Timestamp, PageId, and CustomerId. Each day generates a large log file.

Given two log files, we need to find loyal customers meeting the above criteria.

## Implementation Details

1. `ParseLog`: Reads logs and maps each CustomerId to unique PageIds visited
2. `FindLoyalCustomers`: Identifies customers who visited on both days and viewed at least 2 different pages

## How to Run

Using the Makefile:
```bash
make run
```

With custom log files:
```bash
make run DAY1=path/to/first/log.csv DAY2=path/to/second/log.csv
```

Or directly:
```bash
go run main.go <log_file_day1> <log_file_day2>
```

## Mock Data

Sample log files in CSV format: `Timestamp,PageId,CustomerId`
