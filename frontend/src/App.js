import React, { useState, useEffect } from "react";
import "./App.css";
import {
  Mail,
  // Slack,
  // MessageSquare,
  // MessageCircle,
  // Webhook,
  // Rss,
  // Atom,
  Code,
} from "lucide-react";

const UpdatesDropdown = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [showEmailForm, setShowEmailForm] = useState(false);

  const options = [
    { icon: Mail, label: "Email" },
    // { icon: Slack, label: "Slack" },
    // { icon: MessageSquare, label: "Microsoft Teams" },
    // { icon: MessageCircle, label: "Google Chat" },
    // { icon: Webhook, label: "Webhook" },
    // { icon: Rss, label: "RSS" },
    // { icon: Atom, label: "Atom" },
    { icon: Code, label: "API" },
  ];
  const handleOptionClick = (label) => {
    if (label === "Email") {
      setShowEmailForm(true);
      setIsOpen(false);
    }
  };

  return (
    <div className="updates-container">
      <div
        className="updates-dropdown"
        onMouseEnter={() => setIsOpen(true)}
        onMouseLeave={() => setIsOpen(false)}
      >
        <button
          className="updates-button"
          onMouseEnter={() => setIsOpen(true)}
          onMouseLeave={() => setIsOpen(false)}
        >
          Get updates
        </button>

        {isOpen && (
          <div
            className="dropdown-menu"
            onMouseEnter={() => setIsOpen(true)}
            onMouseLeave={() => setIsOpen(false)}
          >
            <div className="dropdown-content" role="menu">
              {options.map((option, index) => {
                const Icon = option.icon;
                return (
                  <button
                    key={index}
                    className="dropdown-item"
                    role="menuitem"
                    onClick={() => handleOptionClick(option.label)}
                  >
                    <Icon className="dropdown-item-icon" />
                    <span className="dropdown-item-text">{option.label}</span>
                  </button>
                );
              })}
            </div>
          </div>
        )}
      </div>
      {showEmailForm && (
        <EmailSubscriptionForm onClose={() => setShowEmailForm(false)} />
      )}
    </div>
  );
};

const EmailSubscriptionForm = ({ onClose }) => {
  const [emailValue, setEmailValue] = useState("");
  // const [subscriptionType, setSubscriptionType] = useState("all");

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!emailValue || !validateEmail(emailValue)) {
      alert("please enter valid email address");
      return;
    }
    setEmailValue("");
    try {
      const response = await fetch("http://localhost:8080/add-subscriber", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ emailID: emailValue }),
      });
      if (response.ok) {
        alert("successfully subscribed!");
      } else {
        alert("Failed to subscribe.");
      }
      setEmailValue("");
    } catch (error) {
      console.error("Error during subscription:", error);
      alert("An error occured. Please try again.");
    }
  };
  const validateEmail = (email) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  };

  return (
    <div className="modal-overlay">
      <div className="modal-container">
        <div className="modal-header">
          <Mail className="modal-icon" />
          <h2 className="modal-title">Get status updates</h2>
        </div>

        <div className="form-container">
          <div className="form-group">
            <label htmlFor="email" className="form-label">
              Email address
            </label>
            <input
              type="email"
              id="email"
              value={emailValue}
              onChange={(e) => setEmailValue(e.target.value)}
              className="form-input"
              placeholder="you@yourdomain.com"
            />
          </div>

          {/* <div className="radio-group">
            <label className="radio-label">
              <input
                type="radio"
                checked={subscriptionType === "all"}
                onChange={() => setSubscriptionType("all")}
                className="radio-input"
              />
              <span className="radio-text">Get all status updates</span>
            </label>

            <label className="radio-label">
              <input
                type="radio"
                checked={subscriptionType === "specific"}
                onChange={() => setSubscriptionType("specific")}
                className="radio-input"
              />
              <span className="radio-text">Only specific components</span>
            </label>
          </div> */}

          <button className="submit-button" onClick={handleSubmit}>
            Subscribe
          </button>
        </div>

        <button onClick={onClose} className="close-button">
          <svg
            className="close-icon"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>
    </div>
  );
};

// const UptimeBar = ({ daysAgo = 30, uptimeDays = 30 }) => {
//   // Generate array of days with true/false for "up" or "down" status
//   const days = Array.from({ length: daysAgo }, (_, index) => {
//     const isUp = index >= daysAgo - uptimeDays;
//     return isUp;
//   });

//   return (
//     <div className="uptime-bar">
//       <div className="uptime-bar-header">
//         <div className="uptime-bar-service">
//           <div className="uptime-bar-status"></div>
//           <span className="uptime-bar-name">Service Name</span>
//         </div>
//         <span className="uptime-bar-uptime">
//           {((uptimeDays / daysAgo) * 100).toFixed(1)}% uptime
//         </span>
//       </div>

//       <div className="uptime-bar-days">
//         {days.map((isUp, index) => (
//           <div
//             key={index}
//             className={`uptime-bar-day ${
//               isUp ? "uptime-bar-up" : "uptime-bar-down"
//             }`}
//           />
//         ))}
//       </div>

//       <div className="uptime-bar-footer">
//         <span className="uptime-bar-label">{daysAgo} DAYS AGO</span>
//         <span className="uptime-bar-label">TODAY</span>
//       </div>
//     </div>
//   );
// };
const UptimeBar = ({ uptimePercentage }) => {
  const downtimePercentage = 100 - uptimePercentage;

  return (
    <div className="uptime-bar">
      <div className="uptime-bars">
        <svg
          aria-hidden="true"
          preserveAspectRatio="none"
          viewBox="0 0 100 24"
          height="24"
          className="uptime-svg"
        >
          {/* Green uptime bar */}
          <rect
            aria-label={`Uptime: ${uptimePercentage}%`}
            x="0"
            y="0"
            width={uptimePercentage}
            height="24"
            fill="#10b981" // Emerald green for uptime
          ></rect>
          {/* Grey downtime bar */}
          <rect
            aria-label={`Downtime: ${downtimePercentage}%`}
            x={uptimePercentage}
            y="0"
            width={downtimePercentage}
            height="24"
            fill="#d1d5db" // Gray for downtime
          ></rect>
        </svg>
      </div>
      <div className="uptime-labels">
        <span className="uptime-label uptime">{uptimePercentage}% Uptime</span>
        <span className="uptime-label downtime">
          {downtimePercentage}% Downtime
        </span>
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
      <p>Uptime: {uptime}</p>
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
              />
              {/* <UptimeBar uptime={service.uptime} /> */}
              {/* <UptimeBar
                daysAgo={30}
                uptimeDays={28}
                serviceName={service.name}
              /> */}
              <UptimeBar uptimePercentage={service.uptime_percent} />
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
