import React, { useState, useEffect } from "react";
import "./App.css";
import {
  Mail,
  Slack,
  MessageSquare,
  MessageCircle,
  Webhook,
  Rss,
  Atom,
  Code,
} from "lucide-react";

const UpdatesDropdown = () => {
  const [isOpen, setIsOpen] = useState(false);

  const options = [
    { icon: Mail, label: "Email" },
    { icon: Slack, label: "Slack" },
    { icon: MessageSquare, label: "Microsoft Teams" },
    { icon: MessageCircle, label: "Google Chat" },
    { icon: Webhook, label: "Webhook" },
    { icon: Rss, label: "RSS" },
    { icon: Atom, label: "Atom" },
    { icon: Code, label: "API" },
  ];

  return (
    <div className="updates-container">
      <div
        className="updates-dropdown"
        onMouseEnter={() => setIsOpen(true)}
        onMouseLeave={() => setIsOpen(false)}
      >
        <button className="updates-button">Get updates</button>

        {isOpen && (
          <div className="dropdown-menu">
            <div className="dropdown-content" role="menu">
              {options.map((option, index) => {
                const Icon = option.icon;
                return (
                  <button key={index} className="dropdown-item" role="menuitem">
                    <Icon className="dropdown-item-icon" />
                    <span className="dropdown-item-text">{option.label}</span>
                  </button>
                );
              })}
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

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
        <span className={status === "Up" ? "operational" : "down"}>
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
    const interval = setInterval(fetchServiceStatus, 3600);
    return () => clearInterval(interval);
  }, []);
  // const services = [
  //   {
  //     name: "jaypatel.link",
  //     url: "https://jaypatel.link",
  //     status: "Operational",
  //     uptime: 99.9,
  //   },
  //   {
  //     name: "magicdot.jaypatel.link",
  //     url: "https://magicdot.jaypatel.link",
  //     status: "Operational",
  //     uptime: 98.7,
  //   },
  // ];

  // const uptimeData = [
  //   { name: "Jan", uptime: 99.5 },
  //   { name: "Feb", uptime: 99.7 },
  //   { name: "Mar", uptime: 98.9 },
  // ];

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
