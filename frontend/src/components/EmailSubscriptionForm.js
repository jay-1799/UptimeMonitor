import React, { useState } from "react";
import { Mail } from "lucide-react";

const EmailSubscriptionForm = ({ onClose }) => {
  const [emailValue, setEmailValue] = useState("");
  // const [subscriptionType, setSubscriptionType] = useState("all");
  const validateEmail = (email) => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!emailValue || !validateEmail(emailValue)) {
      alert("please enter valid email address");
      return;
    }
    setEmailValue("");
    try {
      const response = await fetch(
        "https://status.jaypatel.link/add-subscriber",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ emailID: emailValue }),
        }
      );
      if (response.ok) {
        alert("Check your inbox for activation link!");
      } else {
        alert("Failed to subscribe.");
      }
      setEmailValue("");
    } catch (error) {
      console.error("Error during subscription:", error);
      alert("An error occured. Please try again.");
    }
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

export default EmailSubscriptionForm;
