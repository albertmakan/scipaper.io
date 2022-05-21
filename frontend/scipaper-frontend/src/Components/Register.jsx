import React from "react";
import * as yup from "yup";
import { register } from "../Services/userService";
import { toast } from "react-toastify";
import { useFormik } from "formik";
import { TextField } from "@mui/material";
import Typography from "@mui/material/Typography";
import { Box } from "@mui/system";
import Button from "@mui/material/Button";
import { Link } from "react-router-dom";

const validationSchema = yup.object({
  email: yup
    .string("Enter your email")
    .email("Enter a valid email")
    .required("Email is required"),
  firstName: yup
    .string("Enter your first name")
    .required("First name is required"),
  lastName: yup
    .string("Enter your last name")
    .required("Last name is required"),
  username: yup.string("Enter your username").required("Username is required"),
  password: yup
    .string("Enter your password")
    .min(8, "Minimum 8 characters required")
    .matches(
      /(?=.*[a-z])/,
      "Password must contain at least 1 lowercase alphabetical character"
    )
    .matches(
      /(?=.*[A-Z])/,
      "Password must contain at least 1 uppercase alphabetical character"
    )
    .matches(
      /(?=.*[0-9])/,
      "Password must contain at least 1 numeric character"
    )
    .required("Required"),
  passwordConfirm: yup
    .string("Confirm your password")
    .required("Confirm your password")
    .oneOf([yup.ref("password"), null], "Passwords must match"),
});

export const Register = () => {
  const formik = useFormik({
    initialValues: {
      firstName: "",
      lastName: "",
      username: "",
      password: "",
      email: "",
      passwordConfirm: "",
    },
    validationSchema,
    onSubmit: (values, { resetForm }) => {
      register(values)
        .then((response) => {
          console.log("response:" + response);
          toast.success("Registration successful! Please log in now.");
          resetForm({});
        })
        .catch((err) => {
          console.log(err);
        });
      resetForm();
    },
  });

  const fieldStyle = {
    paddingTop: "10px",
    paddingBottom: "10px",
    color: "primary",
  };

  const hasError = (fieldName) =>
    formik.touched[fieldName] && Boolean(formik.errors[fieldName]);
  const errorMessage = (fieldName) =>
    formik.touched[fieldName] && formik.errors[fieldName];

  return (
    <Box
      sx={{
        marginLeft: 30,
        marginRight: 30,
        marginTop: 10,
        marginBottom: 0,
      }}
    >
      <Typography variant="h4" align="center">
        Register
      </Typography>
      <Typography component={Link} variant="h6" to="/login">
        Login
      </Typography>
      <form onSubmit={formik.handleSubmit}>
        <TextField
          sx={fieldStyle}
          type="text"
          name="firstName"
          fullWidth
          label="First name"
          value={formik.values.firstName}
          onChange={formik.handleChange}
          error={hasError("firstName")}
          helperText={errorMessage("firstName")}
        />
        <TextField
          sx={fieldStyle}
          type="text"
          name="lastName"
          fullWidth
          label="Last name"
          value={formik.values.lastName}
          onChange={formik.handleChange}
          error={hasError("lastName")}
          helperText={errorMessage("lastName")}
        />
        <TextField
          sx={fieldStyle}
          type="text"
          name="username"
          fullWidth
          label="Username"
          value={formik.values.username}
          onChange={formik.handleChange}
          error={hasError("username")}
          helperText={errorMessage("username")}
        />
        <TextField
          sx={fieldStyle}
          type="email"
          name="email"
          fullWidth
          label="E-mail"
          value={formik.values.email}
          onChange={formik.handleChange}
          error={hasError("email")}
          helperText={errorMessage("email")}
        />
        <TextField
          sx={fieldStyle}
          type="password"
          name="password"
          fullWidth
          label="Password"
          value={formik.values.password}
          onChange={formik.handleChange}
          error={hasError("password")}
          helperText={errorMessage("password")}
        />
        <TextField
          sx={fieldStyle}
          type="password"
          name="passwordConfirm"
          fullWidth
          label="Confirm password"
          value={formik.values.passwordConfirm}
          onChange={formik.handleChange}
          error={hasError("passwordConfirm")}
          helperText={errorMessage("passwordConfirm")}
        />
        <Button variant="contained" type="submit">
          Register
        </Button>
      </form>
    </Box>
  );
};
export default Register;
