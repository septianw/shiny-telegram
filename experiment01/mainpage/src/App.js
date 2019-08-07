import React from 'react';
import Sidebar from './sidebar/Sidebar';
import Breadcrumb from './breadcrumb/Breadcrumb';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { 
  faCog, faSearch, faBell,
  faCalendar, faEnvelope, faAngleDown,
  faUser, faPowerOff, faTasks, faImage } from '@fortawesome/free-solid-svg-icons'
// import logo from './logo.svg';
import './App.css';
import 'bootstrap/dist/css/bootstrap.css';
import './style_v2.css';

function App() {
  return (
    <div id="innerContainer">
      <header className="navbar">
        <div className="container-fluid expanded-panel">
          <div className="row">
            <div id="logo" className="col-xs-12 col-sm-2">
              <a href="index.html">DevOOPS v2</a>
            </div>
            <div id="top-panel" className="col-xs-12 col-sm-10">
              <div className="row">
                <div className="col-xs-8 col-sm-4">
                  <div id="search">
                    <input type="text" placeholder="search" />
                    <FontAwesomeIcon icon={faSearch}/>
                  </div>
                </div>
                <div className="col-xs-4 col-sm-8 top-panel-right">
                  {/* <a href="#" className="about">about</a> */}
                  {/* <a href="index_v1.html" className="style1"></a> */}
                  <ul className="nav navbar-nav pull-right panel-menu">
                    <li className="hidden-xs">
                      <a href="index.html" className="modal-link">
                        <FontAwesomeIcon icon={faBell}/>
                        <span className="badge">7</span>
                      </a>
                    </li>
                    <li className="hidden-xs">
                      <a className="ajax-link" href="ajax/calendar.html">
                        <FontAwesomeIcon icon={faCalendar}/>
                        <i className="fa fa-calendar"></i>
                        <span className="badge">7</span>
                      </a>
                    </li>
                    <li className="hidden-xs">
                      <a href="ajax/page_messages.html" className="ajax-link">
                        <FontAwesomeIcon icon={faEnvelope}/>
                        <span className="badge">7</span>
                      </a>
                    </li>
                    <li className="dropdown">
                      <a href="#" className="dropdown-toggle account" data-toggle="dropdown">
                        <div className="avatar">
                          <img src="img/avatar.jpg" className="img-circle" alt="avatar" />
                        </div>
                        <FontAwesomeIcon icon={faAngleDown} className="pull-right"/>
                        <div className="user-mini pull-right">
                          <span className="welcome">Welcome,</span>
                          <span>Jane Devoops</span>
                        </div>
                      </a>
                      <ul className="dropdown-menu">
                        <li>
                          <a href="#">
                            <FontAwesomeIcon icon={faUser}/>
                            <span>Profile</span>
                          </a>
                        </li>
                        <li>
                          <a href="ajax/page_messages.html" className="ajax-link">
                            <FontAwesomeIcon icon={faEnvelope}/>
                            <span>Messages</span>
                          </a>
                        </li>
                        <li>
                          <a href="ajax/gallery_simple.html" className="ajax-link">
                            <FontAwesomeIcon icon={faImage}/>
                            <span>Albums</span>
                          </a>
                        </li>
                        <li>
                          <a href="ajax/calendar.html" className="ajax-link">
                            <FontAwesomeIcon icon={faTasks}/>
                            <span>Tasks</span>
                          </a>
                        </li>
                        <li>
                          <a href="#">
                            <FontAwesomeIcon icon={faCog}/>
                            <span>Settings</span>
                          </a>
                        </li>
                        <li>
                          <a href="#">
                            <FontAwesomeIcon icon={faPowerOff}/>
                            <span>Logout</span>
                          </a>
                        </li>
                      </ul>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </header>
      <div id="main" className="container-fluid">
        <div className="row">
          <div id="sidebar-left" className="col-xs-2 col-sm-2">
            <Sidebar/>
          </div>
          <div id="content" className="col-xs-12 col-sm-10">
            <Breadcrumb />
            {new Date().toLocaleTimeString()}
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
