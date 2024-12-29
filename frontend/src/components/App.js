import React, { useState, useEffect } from "react";
import "../App.css";

import UpdatesDropdown from "./UpdatesDropdown";
import StatusCard from "./StatusCard";
import IncidentLogs from "./IncidentLogs";
// import UptimeBar from "./UptimeBar";

const App = () => {
  const incidents = [
    { message: "Server downtime", resolved: true },
    { message: "Database issue", resolved: false },
  ];
  const [services, setServices] = useState([]);

  const fetchServiceStatus = async () => {
    try {
      const response = await fetch("http://localhost:8080/status");
      if (response.ok) {
        const data = await response.json();
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
  }, []);

  return (
    <div className="app">
      <header className="app-header">
        <h1>Status Monitoring Dashboard</h1>
      </header>
      <UpdatesDropdown />
      <main className="app-main">
        <div className="status-grid">
          {services.map((service, index) => (
            <div key={index}>
              <StatusCard
                // serviceName={service.name}
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
              {/* <UptimeBar uptimePercentage={service.uptime_percent} /> */}
            </div>
          ))}
        </div>
        <IncidentLogs incidents={incidents} />
      </main>
      <footer className="app-footer">
        <p>&copy; 2024 Jay Patel</p>
      </footer>
    </div>
  );
};

export default App;
