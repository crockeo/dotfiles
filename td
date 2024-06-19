#!/usr/bin/env python3
"""
td -> todo

Implementation of a quick todo app based around graphs.
"""
import argparse
import sqlite3
import sys
from pathlib import Path


def get_db() -> sqlite3.Connection:
    db_dir = Path.home() / ".config"
    db_dir.mkdir(exist_ok=True, parents=True)
    db = sqlite3.connect(db_dir / "td.sqlite3")
    db.execute(
        """
        CREATE TABLE IF NOT EXISTS todo (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL
        )
        """
    )
    db.execute(
        """
        CREATE TABLE IF NOT EXISTS todo_depends_on(
            depends_from INTEGER,
            depends_to INTEGER
        )
        """
    )
    return db


def cmd_new(conn: sqlite3.Connection, args: list[str]) -> None:
    todo_name = " ".join(args)
    conn.execute("INSERT INTO todo (name) VALUES (?)", (todo_name,))
    print(f"Created '{todo_name}'")


def cmd_depends(conn: sqlite3.Connection, args: list[str]) -> None:
    # TODO: make the 1st todo depend on the 2nd todo
    # if none is selected: see how hard it would be to do fzf?
    raise NotImplementedError


def cmd_goals(conn: sqlite3.Connection, args: list[str]) -> None:
    # TODO: show all todos which are not depended on by anything
    raise NotImplementedError


def cmd_action(conn: sqlite3.Connection, args: list[str]) -> None:
    # TODO: show all todos which do not have any dependencies
    raise NotImplementedError


def main(args: list[str]) -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "command",
        choices=[
            "a",
            "action",
            "d",
            "depends",
            "g",
            "goals",
            "n",
            "new",
        ],
    )
    namespace, remaining_args = parser.parse_known_args(args[1:])
    command: str = namespace.command

    with get_db() as conn:
        if command in {"new", "n"}:
            cmd_new(conn, remaining_args)
        elif command in {"depends", "d"}:
            cmd_depends(conn, remaining_args)
        elif command in {"goals", "g"}:
            cmd_goals(conn, remaining_args)
        elif command in {"action", "a"}:
            cmd_action(conn, remaining_args)
        else:
            raise SystemExit(f"Unrecognized command: {command}")
        conn.commit()


if __name__ == "__main__":
    main(sys.argv)
