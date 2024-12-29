import React from "react";
import { FaCheckCircle, FaTimesCircle } from "react-icons/fa";
import UptimeBar from "./UptimeBar";

const StatusCard = ({ serviceName, status, uptime, uptimePercentage }) => {
  return (
    <div className="status-card">
      {/* Service Name and Status Icon */}
      <div className="status-header">
        <h3 className="service-name">
          {status === "Up" ? (
            <FaCheckCircle className="status-icon green" />
          ) : (
            <FaTimesCircle className="status-icon red" />
          )}
          {serviceName}
          <span className="uptime-percentage">{uptimePercentage}% Uptime</span>
        </h3>
      </div>

      {/* Uptime Bar */}
      <UptimeBar uptimePercentage={uptimePercentage} />
    </div>
  );
};

export default StatusCard;
