import './App.css';
import {
    BrowserRouter as Router,
    Route,
    Link, Routes
} from "react-router-dom";
import {useState} from "react";
import axios from "axios";

 const videos = [
{
    ID: 1,
    Name: 'Example Link',
    Link: 'https://example.com',
    Status: true,
    Device: 'Example Device'
}, {
         ID: 2,
         Name: 'Examsple Link',
         Link: 'https://exsample.com',
         Status: false,
         Device: 'Exasmple Device'
     }
]

function App() {
    return (
        <div >
            <HeaderDefault />
        </div>
    );
}

function HeaderDefault() {
    return (
        <Router>
            <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
                <Link to="/" className="navbar-brand">
                    AIFC stream
                </Link>
                <button
                    className="navbar-toggler"
                    type="button"
                    data-toggle="collapse"
                    data-target="#navbarText"
                    aria-controls="navbarText"
                    aria-expanded="false"
                    aria-label="Toggle navigation"
                >
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarText">
                    <ul className="navbar-nav mr-auto">
                        <li className="nav-item active">
                            <Link to="/" className="nav-link">
                                Stream <span className="sr-only">(current)</span>
                            </Link>
                        </li>
                        <li className="nav-item">
                            <Link to="/admin" className="nav-link">
                                Admin page
                            </Link>
                        </li>
                        <li className="nav-item">
                            <Link to="/insert" className="nav-link">
                                Insert
                            </Link>
                        </li>
                    </ul>
                </div>

            </nav>
            <Routes>
                <Route path="/insert" element={<InsertVideo />} />
                <Route path="/admin" element={<AdminPage videos={videos} />} />
                <Route path="/" element={<MainPage devices={videos} />} />
            </Routes>
        </Router>
    );
}


function InsertVideo() {
    return (
        <div className="container m-5">
            <div className="container">
                <h3>How to insert video?</h3>
                <p>Choose video from youtube</p>
                <p>Press share and then embed</p>
                <p>Copy part of text from src within quotes and insert it in form below</p>
                <p>Example:</p>
                <p style={{ display: 'inline' }}>src="</p>
                <p style={{ display: 'inline', backgroundColor: 'yellow' }}>
                    https://www.youtube.com/embed/zIwtfjgAmQM?si=KSQgD6AZEx5g-Vy6
                </p>"
            </div>
            <div className="container">
                <h5>Inserting Form</h5>
                <form method="post">
                    <div className="form-group">
                        <label htmlFor="nameVideo">Name of the video</label>
                        <input
                            type="text"
                            id="nameVideo"
                            name="name"
                            className="form-control"
                            placeholder="Enter name"
                        ></input>
                    </div>
                    <div className="form-group">
                        <label htmlFor="link">Link for the video</label>
                        <input
                            type="text"
                            className="form-control"
                            name="link"
                            id="link"
                            placeholder="Enter Link"
                        ></input>
                    </div>
                    <div className="form-group">
                        <label form="device">Device</label>
                        <input
                            type="text"
                            className="form-control"
                            name="device"
                            id="device"
                            placeholder="Enter Device"
                        ></input>
                    </div>
                    <button type="submit" className="btn btn-primary">
                        Submit
                    </button>
                </form>
            </div>
        </div>
    );
}


function AdminPage({videos}) {
        return (
            <div className="container">
                <div className="d-flex flex-column m-5">
                    {videos.map((video, index) => (
                        <div className="card m-2" style={{ width: '18rem' }} key={index}>
                            <div className="card-body">
                                <h5 className="card-title">{video.Name}</h5>
                                <h6 className="card-subtitle mb-2 text-muted">{video.Device}</h6>

                                {video.Status && <h6 className="card-subtitle mb-2 text-muted">Current video</h6>}

                                {!video.Status && (
                                    <a href={`/admin/${video.Name}`} className="card-link">
                                        Change video
                                    </a>
                                )}

                                <iframe
                                    src={video.Link}
                                    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                                    allowFullScreen
                                    width="100%"
                                    height="60%"
                                ></iframe>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        );
}

function MainPage() {
    const [items, setItems] = useState(null);
    const fetchData = async () => {
        try {
            const response = await axios.get('http://localhost:8080/');
            setItems(response.data);
            console.log(items);
        } catch (error) {
            console.error('Error fetching data:', error.message);
            // Handle the error appropriately (e.g., display a user-friendly message)
        }
    };
    return (
        <div className="container">
            <div className="d-flex flex-column m-5">
                {items.map((device, index) => (
                    <div className="card m-2" style={{ width: '18rem' }} key={index}>
                        <div className="card-body">
                            <h5 className="card-title">{device.Device}</h5>
                            <a href={`/stream/${device.Device}`} className="card-link">
                                Go to the Stream
                            </a>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default App;
