import React, { useState } from "react";
import "./Signup.css";

const Signup = () => {
  const [formData, setFormData] = useState({
    email: "",
    username: "",
    projectLink: "",
  });

  const [validFields, setValidFields] = useState({
    email: false,
    userName: false,
    projectLink: false,
  });

  const validateField = (name, value) => {
    switch (name) {
      case "email":
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value);
      case "username":
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

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("Form submitted:", formData);
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
      </div>
    </div>
  );
};

export default Signup;
