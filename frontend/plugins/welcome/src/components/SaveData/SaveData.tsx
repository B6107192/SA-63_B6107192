import React, { FC,useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
} from '@material-ui/core';
import { makeStyles, } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { EntUser } from '../../api/models/EntUser'; // import interface User
import { EntLocation } from '../../api/models/EntLocation'; // import interface ClubLocation
import { EntClubTypes } from '../../api/models/EntClubTypes'; // import interface ClubType
import { EntActivity} from '../../api/models/EntActivity'; // import interface ClubActivity
import Swal from 'sweetalert2'; // alert
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

interface  Club  {
	CLUBNAME :      string;
	Clubtypes :    number;
	Activity  :    number;
	Location  :    number;
	User      :     number;
}
const SaveData: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [club, setClub] = React.useState<Partial<Club>>({});

  const [users, setUser] = React.useState<EntUser[]>([]);
  const [clublocations, setClubLocation] = React.useState<EntLocation[]>([]);
  const [clubtypes, setClubType] = React.useState<EntClubTypes[]>([]);
  const [clubactivitys, setClubActivity] = React.useState<EntActivity[]>([]);

    // alert setting
    const Toast = Swal.mixin({
      toast: true,
      position: 'top-end',
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
      },
    });

  const getUser = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setUser(res);
  };

  const getClubActivity = async () => {
    const res = await http.listActivity({ limit: 10, offset: 0 });
    setClubActivity(res);
  };

  const getClubLocation = async () => {
    const res = await http.listLocation({ limit: 10, offset: 0 });
    setClubLocation(res);
  };

  const getClubType = async () => {
    const res = await http.listClubTypes({ limit: 10, offset: 0 });
    setClubType(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getUser();
    getClubType();
    getClubLocation();
    getClubActivity();
  }, []);

  // set data to object club
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof SaveData;
    const { value } = event.target;
    setClub({ ...club, [name]: value });
    console.log(club);
  };
  // clear input form
  function clear() {
    setClub({});
  }
  
  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/Clubs';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(club),
    };

    console.log(club); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ระบบจัดการชมรม`}
       subtitle=""
     >
       <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/"
               variant="contained"color="primary"
             >
               ออกจากระบบ
       </Button>
     </Header>
     <Content>
       <ContentHeader title="บันทึกข้อมูล"></ContentHeader>
       <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อชมรม</div>
            </Grid>
            <Grid item xs={9}>
      <form className={classes.root} noValidate autoComplete="off" >
        <TextField label="กรอกชื่อชมรม"name ="CLUBNAME" type="string" value={club.CLUBNAME} onChange={handleChange} className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  />
      </form>
      </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ประเภทชมรม</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>เลือกประเภทชมรม</InputLabel>
        <Select
          name="Clubtypes"
          type="number"
          value={club.Clubtypes || ''}
          onChange={handleChange}
        >
          {clubtypes.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.cLUBETYPENAME}
                      </MenuItem>
                    );
                  })}
        </Select>
      </FormControl>
      </Grid>
      <Grid item xs={3}>
              <div className={classes.paper}>User</div>
            </Grid>
            <Grid item xs={9}>
      <FormControl fullWidth  variant="outlined" 
        className={classes.formControl}>
        <InputLabel>User</InputLabel>
        <Select
          name="User"
          type="number"
          value={club.User || ''}
          onChange={handleChange}
        >
          {users.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.uSERNAME}
                      </MenuItem>
                    );
                  })}
        </Select>
      </FormControl>
      </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>กิจกรรมชมรม</div>
            </Grid>
            <Grid item xs={9}>
      <div>
      <FormControl fullWidth  variant="outlined"
        className={classes.formControl}>
        <InputLabel>เลือกกิจกรรมชมรม</InputLabel>
        <Select
           name="Activity"
           type="number"
           value={club.Activity || ''}
           onChange={handleChange}
         >
           {clubactivitys.map(item => {
                     return (
                       <MenuItem key={item.id} value={item.id}>
                         {item.cLUBEACTIVITYNAME}
                       </MenuItem>
                     );
                   })}
        </Select>
      </FormControl>
      </div>
      </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>สถานที่</div>
            </Grid>
            <Grid item xs={9}>
      <div>
      <FormControl fullWidth  variant="outlined"
        className={classes.formControl}>
        <InputLabel id="demo-simple-select-label">เลือกสถานที่</InputLabel>
        <Select
          name="Location"
          type="number"
          value={club.Location|| ''}
          onChange={handleChange}
        >
          {clublocations.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.cLUBELOCATIONNAME}
                      </MenuItem>
                    );
                  })}
        </Select>
      </FormControl>
      </div>
      </Grid>

          <Grid item xs={3}></Grid>
          <Grid item xs={9}>
             <Button
               onClick={save}
               startIcon={<SaveIcon />}
               variant="contained"
               color="primary"
             >
               บันทึกข้อมูล
             </Button>
           
           </Grid>
          </Grid>
        </Container>
        </Content>
   </Page>
 );
}
export default SaveData;