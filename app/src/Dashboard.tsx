import React from 'react';
import { Title } from 'react-admin';
import { withStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Avatar from '@material-ui/core/Avatar';

const styles = () => ({
  content: {
    display: 'flex',
  },
  avatar: {
    margin: 10,
    width: 400,
    height: 400,
    borderRadius: 0,
  }
})

export default withStyles(styles, { withTheme: true })(({ classes }) => (
    <Card>
        <Title title="快速构建 Golang ❤️ MongoDB ❤️ React 应用" />
        <CardContent className={classes.content}>
          <Avatar
            alt="GoLang"
            src="https://camo.githubusercontent.com/2cc07aaecc587c5eeae78e711a9d71048be9ef41/68747470733a2f2f63646e2d696d616765732d312e6d656469756d2e636f6d2f6d61782f313230302f312a796839306257386a4c346638704f545a5476627a71772e706e67" 
            className={classes.avatar}/>
        </CardContent>
    </Card>
))
