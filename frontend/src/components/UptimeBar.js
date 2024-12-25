import React from "react";

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

export default UptimeBar;
