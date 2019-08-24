#!/usr/bin/python
""" Docker Credentials helper as per
https://docs.docker.com/engine/reference/commandline/login/#credential-helper-protocol """
# Author: Martin Millnert, 2018
# License: GPL-3.0-only
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

import json
import os
import sys

CREDSTORE = os.path.join(os.environ.get("HOME"), ".credhelper")
DB_PATH = os.path.join(CREDSTORE, "creds.json")

def test_create_credstore():
    """ Creates the credstore if it doesn't exist """
    if os.path.exists(CREDSTORE):
        if os.path.isdir(CREDSTORE):
            return True
        else:
            raise Exception('CREDSTORE (%s) occupied by not-a-directory.' % CREDSTORE)
    else:
        os.mkdir(CREDSTORE, 0700)

def read_db():
    """ reads the "db" from CREDSTORE """
    if os.path.isfile(DB_PATH):
        fhd = open(DB_PATH)
        data = json.load(fhd)
        fhd.close()
        return data
    else:
        return {}

def write_db(database):
    """ Writes out database to DB_PATH. """
    test_create_credstore()
    fhd = open(DB_PATH, 'w')
    json.dump(database, fhd)
    fhd.close()

def store(indata):
    """ stores the input data into the DB, appending and overwriting any duplicates. """
    db_content = read_db()
    db_content[indata.get('ServerURL')] = {
        'Username': indata.get('Username'),
        'Secret': indata.get('Secret')
    }
    write_db(db_content)

def get(serverurl):
    """ Returns the credentials for a specific server url"""
    db_content = read_db()
    entry = db_content.get(serverurl)
    if entry:
        sys.stdout.write(json.dumps(entry))

def erase(serverurl):
    """ Erase a server url from the store."""
    db_content = read_db()
    if serverurl in db_content:
        db_content.pop(serverurl)
        write_db(db_content)

def read_stdin_as_json():
    """ Reads stdin, parse as json and return the dict."""
    data = json.load(sys.stdin)
    return data

def read_stdin_as_str():
    """ Reads stdin, parse as str and return."""
    instr = sys.stdin.readline().strip()
    return instr

def main():
    """ main method """
    if len(sys.argv) == 2:
        if sys.argv[1] == "store":
            indata = read_stdin_as_json()
            store(indata)
        elif sys.argv[1] == "erase":
            indata = read_stdin_as_str()
            erase(indata)
        elif sys.argv[1] == "get":
            indata = read_stdin_as_str()
            get(indata)
        else:
            sys.exit(1)
    else:
        sys.exit(1)

if __name__ == '__main__':
    main()
