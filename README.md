# Hazar Ground Station Application

The Hazar Ground Station Application, developed by the Hazar Rocket Team, is designed to monitor real-time data from sensors during rocket flights. The application receives data via serial communication and displays it in a user-friendly interface.

## Data from Development Board

- Rocket speed and position obtained from encoder data
- Rocket rotation obtained from gyroscope data
- Rocket acceleration obtained from accelerometer data
- Atmospheric pressure obtained from barometer data
- Air temperature obtained from temperature sensor data

## Visualization Features of the Application

- Graphs
- Tables
- Numerical values

## Features of the Application

- Real-time data monitoring
- Data recording
- Data analysis

## Usage of the Application

### Processing Data from the Development Board

1. **Data Transmission from Development Board:** The development board sends data from sensors to a computer via serial communication.

2. **Reading Data with USBIPd:** Data received by USBIPd is read by the serial communication driver.
   ![View](images/1.png)

3. **Attaching to WSL with USBIPd:** The desired `busid` is attached to WSL through USBIPd.
   ![View](images/2.png)

4. **Processing Data through WSL:** WSL can process the data by reading the byte array shared by USBIPd.

5. **Processing Data by the Application:** The application can process the data by reading it through WSL.
   ![View](images/4.png)

### Sending to Grafana and Database

After processing the data, the application can save it using one of the following options:

- Grafana: The application can send data to Grafana to visualize it in graphs and tables.
- Database: The application can save data to a database for later analysis.
