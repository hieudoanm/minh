import { addZero } from '@hieudoanm/utils';
import Button from '@mui/material/Button';
import type { NextPage } from 'next';
import Head from 'next/head';
import { useEffect, useState } from 'react';

const APP_NAME = 'POMODORO';

enum Mode {
  REST = 'REST',
  WORK = 'WORK',
}

enum ModeTime {
  REST = 5,
  WORK = 25,
}

enum ModeColor {
  REST = 'bg-blue-500',
  WORK = 'bg-gray-500',
  IDLE = 'bg-white',
}

type Clock = { mode: Mode; status: boolean; seconds: number; display: string };

const PomodoroPage: NextPage = () => {
  const seconds = 1000;

  const [clock, setClock] = useState<Clock>({
    status: false,
    mode: Mode.WORK,
    seconds: ModeTime.WORK * 60,
    display: `${addZero(ModeTime.WORK)}:00`,
  });

  const [timer, setTimer] = useState<NodeJS.Timer | undefined>();

  const start = () => {
    setClock({ ...clock, status: true });
    const timer: NodeJS.Timer = setInterval(() => {
      if (clock.seconds === 0) {
        setClock(({ status, mode }: Clock) => {
          const newMode = mode === Mode.WORK ? Mode.REST : Mode.WORK;
          return {
            status,
            mode: newMode,
            seconds: ModeTime.WORK * 60,
            display: `${addZero(ModeTime.WORK)}:00`,
          };
        });
      } else {
        console.log();
        setClock(({ mode, status, seconds }: Clock) => {
          const secondsLeft = seconds - 1;
          const minutes = addZero(Math.floor(secondsLeft / 60));
          const newSeconds = addZero(secondsLeft % 60);
          return {
            mode,
            status,
            display: `${minutes}:${newSeconds}`,
            seconds: secondsLeft,
          };
        });
      }
    }, seconds);
    setTimer(timer);
  };

  const pause = () => {
    setClock({ ...clock, status: false });
    clearInterval(timer);
  };

  const reset = () => {
    setClock({
      mode: Mode.WORK,
      status: false,
      seconds: ModeTime.WORK * 60,
      display: `${addZero(ModeTime.WORK)}:00`,
    });
    clearInterval(timer);
  };

  useEffect(() => {
    if (clock.seconds === 0) {
      setClock(({ status, mode }: Clock) => {
        const newMode = mode === Mode.WORK ? Mode.REST : Mode.WORK;
        return {
          status,
          mode: newMode,
          seconds: ModeTime[newMode] * 60,
          display: `${addZero(ModeTime[newMode])}:00`,
        };
      });
    }
  }, [clock, timer]);

  useEffect(() => {
    return () => clearInterval(timer);
  }, [timer]);

  const bgColor: ModeColor = clock.status
    ? ModeColor[clock.mode]
    : ModeColor.IDLE;
  const title: string = clock.status
    ? `${APP_NAME} - ${clock.mode}`
    : `${APP_NAME} - IDLE`;

  return (
    <div className={`w-screen h-screen ${bgColor}`}>
      <Head>
        <title>{title}</title>
      </Head>
      <div className="container h-full mx-auto p-8">
        <div className="w-full h-full flex items-center justify-center">
          <div className="rounded p-8 border shadow-2xl bg-white">
            <h1 className="text-center text-4xl sm:text-5xl uppercase">
              {APP_NAME}
            </h1>
            <div className="py-8">
              <p className="text-center text-8xl sm:text-9xl">
                {clock.display}
              </p>
            </div>
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-8">
              <div>
                <Button
                  onClick={clock.status ? pause : start}
                  variant="outlined"
                  className="w-full uppercase"
                >
                  {clock.status
                    ? `${
                        clock.mode === Mode.WORK
                          ? `${Mode.WORK}`
                          : `${Mode.REST}`
                      }ING`
                    : `Start`}
                </Button>
              </div>
              <div>
                <Button
                  onClick={reset}
                  variant="outlined"
                  className="w-full uppercase"
                >
                  Reset
                </Button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default PomodoroPage;
