import './App.css';
import {
    BrowserRouter as Router,
    Route,
    Link, Routes,useParams
} from "react-router-dom";
import {useEffect, useState} from "react";
import axios from "axios";

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
                <Route path="/admin" element={<AdminPage />} />
                <Route path="/" element={<MainPage  />} />
                <Route path="/stream/:id" element={<StreamPage />} />
            </Routes>
        </Router>
    );
}


const InsertVideo = () => {
    const [name, setName] = useState('');
    const [link, setLink] = useState('');
    const [device, setDevice] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            const response = await fetch('http://localhost:8080/insert', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: new URLSearchParams({
                    name: name,
                    link: link,
                    device: device,
                }),
            });

            const data = await response.json();
            console.log(data);

            // Optionally, you can reset the form fields after a successful submission
            setName('');
            setLink('');
            setDevice('');
        } catch (error) {
            console.error('Error submitting data:', error.message);
        }
    };

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
                <form onSubmit={handleSubmit}>
                    <div className="form-group">
                        <label htmlFor="nameVideo">Name of the video</label>
                        <input
                            type="text"
                            id="nameVideo"
                            name="name"
                            className="form-control"
                            placeholder="Enter name"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                        />
                    </div>
                    <div className="form-group">
                        <label htmlFor="link">Link for the video</label>
                        <input
                            type="text"
                            className="form-control"
                            name="link"
                            id="link"
                            placeholder="Enter Link"
                            value={link}
                            onChange={(e) => setLink(e.target.value)}
                        />
                    </div>
                    <div className="form-group">
                        <label form="device">Device</label>
                        <input
                            type="text"
                            className="form-control"
                            name="device"
                            id="device"
                            placeholder="Enter Device"
                            value={device}
                            onChange={(e) => setDevice(e.target.value)}
                        />
                    </div>
                    <button type="submit" className="btn btn-primary">
                        Submit
                    </button>
                </form>
            </div>
        </div>
    );
};


function AdminPage() {
    const [videos, setVideos] = useState(null);

    const fetchData = async () => {
        try {
            const response = await axios.get('http://localhost:8080/admin');
            setVideos(response.data);
            console.log(response.data)
        } catch (error) {
            console.error('Error fetching data:', error.message);
            // Handle the error appropriately (e.g., display a user-friendly message)
        }
    };

    useEffect(() => {
        const intervalId = setInterval(() => {
            fetchData();
        }, 1000); // Fetch data every 60 seconds

        return () => clearInterval(intervalId); // Cleanup the interval on component unmount
    }, []);



        return (
            <div className="container">
                <div className="d-flex flex-column m-5">
                    {videos ? (
                        videos.map((video, index) => (
                                <div className="card m-2" style={{ width: '18rem' }} key={index}>
                                    <div className="card-body">
                                        <h5 className="card-title">{video.name}</h5>
                                        <h6 className="card-subtitle mb-2 text-muted">{video.device}</h6>

                                        {video.status && <h6 className="card-subtitle mb-2 text-muted">Current video</h6>}

                                        {!video.status && (
                                            <a href={`/admin/${video.name}`} className="card-link">
                                                Change video
                                            </a>
                                        )}

                                        <iframe
                                            src={video.link}
                                            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                                            allowFullScreen
                                            width="100%"
                                            height="60%"
                                        ></iframe>
                                    </div>
                                </div>
                            ))
                    ):(
                        <p>Loading...</p>
                    )}

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
            console.log(response.data)
        } catch (error) {
            console.error('Error fetching data:', error.message);
            // Handle the error appropriately (e.g., display a user-friendly message)
        }
    };

    useEffect(() => {
        const intervalId = setInterval(() => {
            fetchData();
        }, 1000); // Fetch data every 60 seconds

        return () => clearInterval(intervalId); // Cleanup the interval on component unmount
    }, []);


    return (
        <div className="container">
            {items ? (
                <div className="d-flex flex-column m-5">
                    {items.map((device, index) => (
                        <div className="card m-2" style={{ width: '18rem' }} key={index}>
                            <div className="card-body">
                                <h5 className="card-title">{device.device}</h5>
                                <a href={`/stream/${device.device}`} className="card-link">
                                    Go to the Stream
                                </a>
                            </div>
                        </div>
                    ))}
                </div>
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
}

const StreamPage = () => {
    const { id } = useParams();
    const [link, setLink] = useState(null);

    useEffect(() => {
        const fetchStreamData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/stream/${id}`);
                const data = await response.json();
                setLink(data);
            } catch (error) {
                console.error('Error fetching stream data:', error.message);
            }
        };

        fetchStreamData();
    }, [id]);

    return (
        <div style={{height:"100vh"}}>
            {link ? (

                <iframe
                    src={link.link}
                    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                    allowFullScreen
                    width="100%"
                    height="100%"
                ></iframe>


            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};


export default App;
