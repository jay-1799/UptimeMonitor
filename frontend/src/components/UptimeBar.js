import React from "react";

const UptimeBar = ({ uptimePercentage }) => {
  // const downtimePercentage = 100 - uptimePercentage;
  const barHeight = 38;
  const rectWidth = 10;
  const gap = 1;
  const totalDays = 45;

  const activeDays = Math.floor((uptimePercentage * totalDays) / 100);
  return (
    <svg
      width={(rectWidth + gap) * totalDays}
      height={barHeight}
      className="uptime-bar-svg"
    >
      {Array.from({ length: totalDays }).map((_, index) => (
        <rect
          key={index}
          // className="uptime-bar-rect"
          className={
            index < activeDays
              ? "uptime-bar-rect active"
              : "uptime-bar-rect inactive"
          }
          x={index * (rectWidth + gap)}
          y="2"
          width={rectWidth}
          height="24"
          // fill={index < activeDays ? "#14D433" : "#A3A3A3"} // Green for active days, gray for others
          rx="1"
        />
      ))}
    </svg>
  );
};

export default UptimeBar;
