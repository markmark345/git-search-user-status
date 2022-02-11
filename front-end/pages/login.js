import Head from "next/head";
import Image from "next/image";
import { ContextProvider, GlobalContext } from "../context/ContextProvider ";
import {
  Grid,
  Paper,
  Avatar,
  TextField,
  Button,
  Typography,
  Link,
  FormControlLabel,
  Checkbox,
} from "@mui/material";
import GitHubIcon from "@mui/icons-material/GitHub";

const UserName = () => {
  const [user, setUser] = useContext(GlobalContext)
  const btnstyle = { margin: "8px 0" };
  return (
    <Grid>
      <Paper elevation={10} sx={{ p: 8, width: 240, height: 280, margin:"20px auto" }}>
        <Grid align="center">
          <Avatar sx={{ bgcolor: "#000000" }}>
            <GitHubIcon />
          </Avatar>
          <h2>Git Username</h2>
        </Grid>
        <TextField
          label="Username"
          placeholder="Enter username"
          fullWidth
          required
        />
        <Button
          type="submit"
          color="primary"
          variant="contained"
          style={btnstyle}
          fullWidth
          onClick={() => setUser()}
        >
          Sign in
        </Button>
      </Paper>
    </Grid>
  );
};

export default UserName;
