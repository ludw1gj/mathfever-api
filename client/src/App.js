import React, { Component } from 'react';
import { BrowserRouter } from 'react-router-dom';

import AppRouter from './router';

import Navbar from './components/Navbar';
import Drawer from './components/Drawer';
import Footer from './components/Footer';
import MDLLayout from './generic/MDLLayout';
import MDLContent from './generic/MDLContent';
import Content from './generic/Content';

/** The main component. */
class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div className="App">
          <MDLLayout>
            <Navbar />
            <Drawer />

            <MDLContent>
              <Content>
                <AppRouter />
              </Content>
              <Footer />
            </MDLContent>
          </MDLLayout>
        </div>
      </BrowserRouter>
    );
  }
}

export default App;
