// Static Frontend for Status Monitoring Dashboard with Dark Theme and Enhanced Features

import React from "react";
import "./App.css";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
} from "recharts";

const UptimeBar = ({ uptime }) => {
  const uptimeColor =
    uptime > 98 ? "#4caf50" : uptime > 95 ? "#ffc107" : "#f44336";

  return (
    <div className="uptime-bar">
      <div
        className="uptime-fill"
        style={{ width: `${uptime}%`, backgroundColor: uptimeColor }}
      >
        {uptime}%
      </div>
    </div>
  );
};

const StatusCard = ({ serviceName, status, uptime }) => {
  return (
    <div className="status-card">
      <h3>{serviceName}</h3>
      <p>
        Status:{" "}
        <span className={status === "Operational" ? "operational" : "down"}>
          {status}
        </span>
      </p>
      <p>Uptime: {uptime}%</p>
    </div>
  );
};

const IncidentLogs = ({ incidents }) => {
  return (
    <div className="incident-logs">
      <h3>Incident Logs</h3>
      <ul>
        {incidents.map((incident, index) => (
          <li
            key={index}
            className={incident.resolved ? "resolved" : "unresolved"}
          >
            {incident.message} - {incident.resolved ? "Resolved" : "Unresolved"}
          </li>
        ))}
      </ul>
    </div>
  );
};

const App = () => {
  const services = [
    {
      name: "jaypatel.link",
      url: "https://jaypatel.link",
      status: "Operational",
      uptime: 99.9,
    },
    {
      name: "magicdot.jaypatel.link",
      url: "https://magicdot.jaypatel.link",
      status: "Operational",
      uptime: 98.7,
    },
  ];

  const incidents = [
    { message: "Server downtime", resolved: true },
    { message: "Database issue", resolved: false },
  ];

  const uptimeData = [
    { name: "Jan", uptime: 99.5 },
    { name: "Feb", uptime: 99.7 },
    { name: "Mar", uptime: 98.9 },
  ];

  return (
    <div className="app">
      <header className="app-header">
        <h1>Status Monitoring Dashboard</h1>
      </header>
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
              />
              <UptimeBar uptime={service.uptime} />
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
