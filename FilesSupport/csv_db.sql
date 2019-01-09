-- phpMyAdmin SQL Dump
-- version 3.3.9
-- http://www.phpmyadmin.net
--
-- Servidor: localhost
-- Tiempo de generación: 10-12-2018 a las 19:18:01
-- Versión del servidor: 5.5.8
-- Versión de PHP: 5.3.5

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Base de datos: `csv_db`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `tbl_name`
--

CREATE TABLE IF NOT EXISTS `tbl_name` (
  `rowid` int(1) DEFAULT NULL,
  `userid` varchar(11) DEFAULT NULL,
  `userstate` varchar(6) DEFAULT NULL,
  `usermanager` varchar(34) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Volcar la base de datos para la tabla `tbl_name`
--

INSERT INTO `tbl_name` (`rowid`, `userid`, `userstate`, `usermanager`) VALUES
(1, 'etrossero', 'activo', 'romina.marzovilla@mercadolibre.com'),
(2, 'dbertoni', 'activo', 'romina.marzovilla@mercadolibre.com'),
(3, 'rbochini', 'activo', 'ernesto.vincelli@mercadolibre.com'),
(4, 'dgarnero', 'activo', 'romina.marzovilla@mercadolibre.com'),
(5, 'jopastoriza', 'activo', 'ernesto.vincelli@mercadolibre.com'),
(6, 'laislas', 'activo', 'ernesto.vincelli@mercadolibre.com');
