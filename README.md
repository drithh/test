# Technical Test: Dana Bagus

## Project Setup and Execution

To run this project, please follow these steps:

1. **Clone the Repository**  
   Begin by cloning this repository to your local machine.

2. **Run the Application**  
   Navigate to each folder and execute the command:
   ```bash
   go run .
   ```

## Foobar

### Overview

This is a simple foobar test, similar to the classic FizzBuzz challenge. The results are displayed in the following format:

```
Bar, Foo, 98, Foo, Bar, 94, Foo, 92, 91, FooBar, 88, Foo, 86, Bar, Foo, 82, Foo, Bar, Foo, 77, 76, FooBar, 74, Foo, Bar, Foo, 68, Foo, Bar, 64, Foo, 62, FooBar, 58, Foo, 56, Bar, Foo, 52, Foo, Bar, 49, Foo, 46, FooBar, 44, Foo, Bar, Foo, 38, Foo, Bar, 34, Foo, 32, FooBar, 28, Foo, 26, Bar, Foo, 22, Foo, Bar, Foo, 16, FooBar, 14, Foo, Bar, Foo, 8, Foo, 4, 1
```

## Weather Forecast

### Overview

This project provides a weather forecast for the next 5 days, utilizing data from the OpenWeatherMap API. The results are displayed as follows:

```
5-Day Weather Forecast for Jakarta (Mode: Midday)
Sun 06 Oct 2024: 32.07°C
Mon 07 Oct 2024: 31.51°C
Tue 08 Oct 2024: 32.97°C
Wed 09 Oct 2024: 30.78°C
Thu 10 Oct 2024: 32.91°C
```

### Weather Forecast Modes

This project supports three distinct forecasting modes for the next 5 days:

1. **Start Day**  
   Provides a forecast based on a reference day, which can be set between 0 and 3.

2. **Mid Day**  
   Generates a forecast centered around the middle of the day, with reference points from 11 to 13.

3. **Average**  
   Offers a forecast based on the average temperature over the specified period.
