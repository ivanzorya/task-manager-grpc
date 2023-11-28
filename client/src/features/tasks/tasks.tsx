import React, { useState, useEffect } from "react";
import { Metadata } from "nice-grpc-web";
import { Button, Form, Container, Modal, InputGroup } from "react-bootstrap"
import TaskRow from "./single-task"
import { clientTasks } from "../../client/client"
import { Task } from "../../app/proto/tasks";
import { accessToken, emailLogged } from "../../App";
import { useRecoilState } from "recoil";

type TasksProps = {
    setAlertMessage: (val: string) => void,
    setAlertVariant: (val: string) => void,
    setIsAlertVisible: (val: boolean) => void,
}

export const Tasks = ({ setAlertMessage, setAlertVariant, setIsAlertVisible }: TasksProps) => {
    const [token] = useRecoilState(accessToken);
    const [emailRecoiled] = useRecoilState(emailLogged);

    const [tasks, setTasks] = useState<Task[]>([]);

    const [addNewTask, setAddNewTask] = useState(false);
    const [newTask, setNewTask] = useState({ "subject": "", "done": false });

    const changeSingleTask = (updatedTask: Task) => {
        const task = {
            task: updatedTask,
        }
        clientTasks.queryUpdateTask(
            task,
            {
                metadata: Metadata({ 'Authorization': token }),
            },
        ).then(() => getAllTasks()).catch((e) => {
            setAlertMessage("Task updating fail");
            setAlertVariant("danger");
            setIsAlertVisible(true);
            getAllTasks();
        })
    }

    const addSingleTask = () => {
        setAddNewTask(false);
        const task = {
            task: {
                id: "fake",
                subject: newTask.subject,
                done: newTask.done,
                userEmail: emailRecoiled,
            }
        }
        clientTasks.queryCreateTask(
            task,
            {
                metadata: Metadata({ 'Authorization': token })
            },
        ).then(() => getAllTasks()).catch((e) => {
            setAlertMessage("Task creation fail");
            setAlertVariant("danger");
            setIsAlertVisible(true);
        })
    }

    const getAllTasks = () => clientTasks.queryGetTasks({}, {
        metadata: Metadata({ 'Authorization': token }),
    }).then(res => setTasks(res.tasks));

    const deleteSingleTask = (id: string) => {
        clientTasks.queryDeleteTask({ id: id }).then(() => getAllTasks());
    }

    useEffect(() => {
        if (token !== "") {
            getAllTasks();
        }
        setTasks([]);
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [token])


    return (token !== "" ?
        <div>

            <Container >
                <Button className="stndrt-class" onClick={() => setAddNewTask(true)}>Add new task</Button>
            </Container>

            <Container>
                {tasks != null && tasks.map((task) => (
                    <TaskRow key={task.id} taskData={task}
                        deleteSingleTask={deleteSingleTask} changeSingleTask={changeSingleTask} />
                ))}
            </Container>

            <Modal show={addNewTask} onHide={() => setAddNewTask(false)} centered>
                <Modal.Header closeButton>
                    <Modal.Title className="stndrt-class">Add Task</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <Form.Group >
                        <Form.Control className="stndrt-class" onChange={(event) => {
                            setNewTask({ ...newTask, subject: event.target.value })
                        }} />
                        <InputGroup.Checkbox type="checkbox" onChange={
                            (event: React.ChangeEvent<HTMLInputElement>) => {
                                setNewTask({ ...newTask, done: event.target.checked })
                            }}
                        />
                    </Form.Group>
                    <Button className="stndrt-class" onClick={() => addSingleTask()}>Add</Button>
                    <Button className="stndrt-class" onClick={() => setAddNewTask(false)}>Cancel</Button>
                </Modal.Body>
            </Modal>
        </div>
        :
        null
    );
}

export default Tasks