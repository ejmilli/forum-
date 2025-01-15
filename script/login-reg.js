"use strict";

// Show Sign Up Form when 'Sign Up' is clicked
document
  .getElementById("showSignUpForm")
  .addEventListener("click", function (e) {
    e.preventDefault();
    document.getElementById("signInForm").style.display = "none"; // Hide login form
    document.getElementById("signUpForm").style.display = "block"; // Show sign-up form
  });

// Show Sign In Form when 'Sign In' is clicked
document
  .getElementById("showSignInForm")
  .addEventListener("click", function (e) {
    e.preventDefault();
    document.getElementById("signUpForm").style.display = "none"; // Hide sign-up form
    document.getElementById("signInForm").style.display = "block"; // Show login form
  });
