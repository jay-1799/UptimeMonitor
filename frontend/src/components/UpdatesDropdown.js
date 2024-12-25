import React, { useState } from "react";
import EmailSubscriptionForm from "./EmailSubscriptionForm";
import { Mail, Code } from "lucide-react";

const UpdatesDropdown = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [showEmailForm, setShowEmailForm] = useState(false);

  const options = [
    { icon: Mail, label: "Email" },
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

export default UpdatesDropdown;
