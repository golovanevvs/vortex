-- Master initialization script for PostgreSQL replication
-- This script runs automatically when the master container starts for the first time

-- Create application database if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'delayed_notifier') THEN
        CREATE DATABASE delayed_notifier;
        RAISE NOTICE 'Database delayed_notifier created';
    ELSE
        RAISE NOTICE 'Database delayed_notifier already exists';
    END IF;
END
$$;

-- Switch to application database
\c delayed_notifier

-- Create replication user if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'repl_user') THEN
        CREATE USER repl_user WITH REPLICATION ENCRYPTED PASSWORD 'repl_password';
        RAISE NOTICE 'Replication user repl_user created';
    ELSE
        RAISE NOTICE 'Replication user repl_user already exists';
    END IF;
END
$$;

-- Create replication slots if they don't exist
DO $$
BEGIN
    -- Slot for slave 1
    IF NOT EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'replication_slot_slave1') THEN
        PERFORM pg_create_physical_replication_slot('replication_slot_slave1', true);
        RAISE NOTICE 'Replication slot replication_slot_slave1 created';
    ELSE
        RAISE NOTICE 'Replication slot replication_slot_slave1 already exists';
    END IF;

    -- Slot for slave 2  
    IF NOT EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'replication_slot_slave2') THEN
        PERFORM pg_create_physical_replication_slot('replication_slot_slave2', true);
        RAISE NOTICE 'Replication slot replication_slot_slave2 created';
    ELSE
        RAISE NOTICE 'Replication slot replication_slot_slave2 already exists';
    END IF;
END
$$;

-- Configure replication access in pg_hba.conf (this will be appended by the setup script)
-- Note: pg_hba.conf modifications are done by the setup-replication.sh script

-- Verify setup
DO $$
BEGIN
    RAISE NOTICE '=== Master Initialization Complete ===';
    RAISE NOTICE 'Database: delayed_notifier';
    RAISE NOTICE 'Replication user: repl_user';
    RAISE NOTICE 'Replication slots: replication_slot_slave1, replication_slot_slave2';
END
$$;

-- Show current status
SELECT 
    (SELECT count(*) FROM pg_database WHERE datname = 'delayed_notifier') as db_exists,
    (SELECT count(*) FROM pg_roles WHERE rolname = 'repl_user') as repl_user_exists,
    (SELECT count(*) FROM pg_replication_slots WHERE slot_name IN ('replication_slot_slave1', 'replication_slot_slave2')) as slots_created;