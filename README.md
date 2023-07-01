# Groupie Trackers

Groupie Trackers is a website that displays information about bands and artists, their concert locations, and dates. The project aims to create a user-friendly site with various data visualizations to present the data in an organized and visually appealing manner.

## Features

### Band Information
The website provides detailed information about bands and artists, including their names, images, the year they started their activity, the date of their first album, and the members. The data is fetched from the API's `artists` endpoint.

### Concert Locations
Concert locations are displayed, showcasing the last and/or upcoming venues where the bands will perform. The data is retrieved from the API's `locations` endpoint.

### Concert Dates
The website also presents the dates of the last and/or upcoming concerts of the bands. This information is obtained from the API's `dates` endpoint.

### Data Visualization
The bands' information and concert details are visualized using various data visualization techniques, such as blocks, cards, tables, lists, pages, and graphics. The choice of visualization method is left to the developer's discretion, with a focus on user-friendly and visually appealing displays.

### Client-Server Communication
The website features client-server communication, allowing users to trigger actions that interact with the server. This event-based system enables the client to send requests to the server, which responds with the requested information following the request-response model.

### Error Handling
The website and server are designed to be robust and should not crash under any circumstances. All pages are thoroughly tested to ensure they work correctly, and appropriate error handling mechanisms are implemented to gracefully handle any errors.

## Technologies Used

- Backend: Go programming language (Golang)
- Frontend: HTML, CSS, JavaScript (or any preferred frontend framework)

## Getting Started

To set up and run the Groupie Trackers project, follow these steps:

1. Clone the repository:

``` shel
    git clone https://github.com/alpapie/groupie-tracker-.git
```

2. Run the project:

``` bash
    go run . --port=<PORT>
```
3. Open the website in your browser:
 ```
  http://localhost:<PORT>
```

## License & Copyright
**[abdouksow](https://learn.zone01dakar.sn/git/abdouksow)** <br>
**[mouhametadiouf](https://learn.zone01dakar.sn/git/mouhametadiouf)**<br>
**[mamoundiaye](https://learn.zone01dakar.sn/git/mamoundiaye)**