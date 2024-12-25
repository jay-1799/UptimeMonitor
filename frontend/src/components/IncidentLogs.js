import React from "react";

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

export default IncidentLogs;
