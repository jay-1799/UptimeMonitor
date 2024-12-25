import React from "react";

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

export default StatusCard;
