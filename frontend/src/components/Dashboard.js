import React, { useState, useEffect } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
  useParams,
} from "react-router-dom";
import "../App.css";
import Login from "./Login";
import UpdatesDropdown from "./UpdatesDropdown";
import StatusCard from "./StatusCard";
import IncidentLogs from "./IncidentLogs";

const Dashboard = () => {
  const { username } = useParams(); // Get the username from the route
  const [services, setServices] = useState([]);

  const fetchServiceStatus = async () => {
    try {
      const response = await fetch(
        // `https://status.jaypatel.link/status`
        `http://localhost:8080/get-status-page?username=${username}`
      );
      if (response.ok) {
        const data = await response.json();
        console.log(data);
        setServices(data);
      } else {
        console.error("Failed to fetch service statuses");
      }
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  useEffect(() => {
    fetchServiceStatus();
    const interval = setInterval(fetchServiceStatus, 3600000);
    return () => clearInterval(interval);
  }, [username]);

  return (
    <div className="app">
      <header className="app-header">
        <h1>{username}'s Status Monitoring Dashboard</h1>
      </header>
      <UpdatesDropdown />
      <main className="app-main">
        <div className="status-grid">
          {services.map((service, index) => (
            <div key={index}>
              <StatusCard
                serviceName={
                  <a
                    href={service.url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="service-link"
                  >
                    {service.name}
                  </a>
                }
                status={service.status}
                uptime={service.uptime}
                uptimePercentage={service.uptime_percent}
              />
            </div>
          ))}
        </div>
        <IncidentLogs incidents={services.incidents || []} />
      </main>
      <footer className="app-footer">
        <p>&copy; 2024 Jay Patel</p>
      </footer>
    </div>
  );
};

export default Dashboard;
