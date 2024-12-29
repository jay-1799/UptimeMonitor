import React from "react";

const UptimeBar = ({ uptimePercentage }) => {
  // const downtimePercentage = 100 - uptimePercentage;
  const barHeight = 38;
  const rectWidth = 10;
  const gap = 1;
  const totalDays = 60;

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

  // return (
  //   <div className="uptime-bar">
  //     <div className="uptime-bars">
  //       <svg
  //         aria-hidden="true"
  //         preserveAspectRatio="none"
  //         viewBox="0 0 100 24"
  //         height="24"
  //         className="uptime-svg"
  //       >
  //         {/* Green uptime bar */}
  //         <rect
  //           aria-label={`Uptime: ${uptimePercentage}%`}
  //           x="0"
  //           y="0"
  //           width={uptimePercentage}
  //           height="24"
  //           fill="#10b981" // Emerald green for uptime
  //         ></rect>
  //         {/* Grey downtime bar */}
  //         <rect
  //           aria-label={`Downtime: ${downtimePercentage}%`}
  //           x={uptimePercentage}
  //           y="0"
  //           width={downtimePercentage}
  //           height="24"
  //           fill="#d1d5db" // Gray for downtime
  //         ></rect>
  //       </svg>
  //     </div>
  //     <div className="uptime-labels">
  //       <span className="uptime-label uptime">{uptimePercentage}% Uptime</span>
  //       <span className="uptime-label downtime">
  //         {downtimePercentage}% Downtime
  //       </span>
  //     </div>
  //   </div>
  // );
};

export default UptimeBar;
