import React, { useState } from "react";
import "./Signup.css";

const Signup = () => {
  const [formData, setFormData] = useState({
    email: "",
    userName: "",
    projectLink: "",
  });

  const [validFields, setValidFields] = useState({
    email: false,
    userName: false,
    projectLink: false,
  });

  const [statusMessage, setStatusMessage] = useState("");

  const validateField = (name, value) => {
    switch (name) {
      case "email":
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value);
      case "userName":
        return value.length >= 2;
      case "projectLink":
        try {
          const url = new URL(value);
          return url.protocol === "http:" || url.protocol === "https:";
        } catch {
          return false;
        }
      default:
        return false;
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));

    setValidFields((prev) => ({
      ...prev,
      [name]: validateField(name, value),
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Check if all fields are valid before submitting
    if (
      !validFields.email ||
      !validFields.userName ||
      !validFields.projectLink
    ) {
      setStatusMessage("Please ensure all fields are valid before submitting.");
      return;
    }

    try {
      const response = await fetch("http://localhost:8080/create-status-page", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          emailID: formData.email,
          username: formData.userName,
          projectLink: formData.projectLink,
        }),
      });

      if (response.ok) {
        setStatusMessage("Status page created successfully!");
        setFormData({
          email: "",
          userName: "",
          projectLink: "",
        });
        setValidFields({
          email: false,
          userName: false,
          projectLink: false,
        });
      } else {
        setStatusMessage("Failed to create status page. Please try again.");
      }
    } catch (error) {
      console.error("Error:", error);
      setStatusMessage("An error occurred. Please try again later.");
    }
  };

  return (
    <div className="page">
      <div className="signup-container">
        <form onSubmit={handleSubmit}>
          <div className={`form-group ${validFields.email ? "valid" : ""}`}>
            <label htmlFor="email">Email address</label>
            <input
              type="email"
              id="email"
              name="email"
              value={formData.email}
              onChange={handleChange}
              className="form-input"
              placeholder="email@company.com"
            />
          </div>

          <div className={`form-group ${validFields.userName ? "valid" : ""}`}>
            <label htmlFor="userName">Username</label>
            <input
              type="text"
              id="userName"
              name="userName"
              value={formData.userName}
              onChange={handleChange}
              className="form-input"
              placeholder="Your username"
            />
          </div>

          <div
            className={`form-group ${validFields.projectLink ? "valid" : ""}`}
          >
            <label htmlFor="projectLink">Project Link</label>
            <input
              type="text"
              id="projectLink"
              name="projectLink"
              value={formData.projectLink}
              onChange={handleChange}
              className="form-input"
              placeholder="Project URL"
            />
          </div>

          <button type="submit" className="submit-button">
            GET YOUR STATUS PAGE
          </button>
        </form>
        {statusMessage && <p className="status-message">{statusMessage}</p>}
      </div>
    </div>
  );
};

export default Signup;
